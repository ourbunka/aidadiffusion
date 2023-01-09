import random
from typing import List, Optional, Union

from PIL import Image
import os
import sys
# Add vendor directory to module search path
parent_dir = os.path.abspath(os.path.dirname(__file__))
vendor_dir = os.path.join(parent_dir, 'Vendor')

sys.path.append(vendor_dir)
import torch  #pip3 install torch torchvision torchaudio --extra-index-url https://download.pytorch.org/whl/cu117 and then copy to vendor directory

#pip install xformers==0.0.16rc402    install xformers

from diffusers.models import AutoencoderKL, UNet2DConditionModel
from diffusers.pipeline_utils import DiffusionPipeline
from diffusers.schedulers import DDIMScheduler, LMSDiscreteScheduler, PNDMScheduler
from diffusers import StableDiffusionUpscalePipeline

def cudasdfunc(promtInput, iterationsNum, saveName, inputwidth, inputheight ): #add upscalesetting
    #torch.backends.cudnn.benchmark = True
    torch.backends.cuda.matmul.allow_tf32 = True

    seed = random.randint(0,9999999999999)
    print(seed)
    print(torch.__version__)
    
    
    schedulersetting = LMSDiscreteScheduler.from_pretrained("./stable-diffusion-2-1", subfolder="scheduler", revision="fp16",torch_dtype=torch.float16)
    pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-2-1", scheduler=schedulersetting).to('cuda') 
    torch.manual_seed(seed)


    prompt = promtInput
    image = pipe(prompt, height=inputheight, width=inputwidth, num_inference_steps=iterationsNum, guidance_scale=7.5, eta=0.0).images[0]
    
    
    upscaleSetting = 0 #todo replace with upscaleSetting = upscalesetting, need to implement in go api and ui

    if upscaleSetting == 1:
        image = image.resize((192, 192))
        upscalePipeline = StableDiffusionUpscalePipeline.from_pretrained("./stable-diffusion-x4-upscaler", torch_dtype=torch.float16).to("cuda") #need at least 24gb of vram for 192px input, 768px upscaled output.
        upscalePipeline.enable_xformers_memory_efficient_attention
        image = upscalePipeline(prompt="", image=image).images[0]
        upscaleSaveName = saveName + "upscaled.png"
        image.save(upscaleSaveName)

    if upscaleSetting == 0:
        image.save(saveName)
    
   