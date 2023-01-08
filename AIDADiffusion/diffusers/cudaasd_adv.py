import random
from typing import List, Optional, Union


import os
import sys
# Add vendor directory to module search path
parent_dir = os.path.abspath(os.path.dirname(__file__))
vendor_dir = os.path.join(parent_dir, 'Vendor')

sys.path.append(vendor_dir)
import torch  #pip3 install torch torchvision torchaudio --extra-index-url https://download.pytorch.org/whl/cu117 and then copy to vendor directory


from diffusers.models import AutoencoderKL, UNet2DConditionModel
from diffusers.pipeline_utils import DiffusionPipeline
from diffusers.schedulers import DDIMScheduler, LMSDiscreteScheduler, PNDMScheduler

def cudaadvfunc(promtInput, iterationsNum, saveName, userSeedInput, userInputScheduler, userInputGuidance, inputwidth, inputheight ):
    torch.backends.cuda.matmul.allow_tf32 = True
    seed = userSeedInput
    print(seed)

    #scheduler/samplers

    if userInputScheduler == "lms":
        schedulersetting = LMSDiscreteScheduler.from_pretrained("./stable-diffusion-2-1", subfolder="scheduler", revision="fp16",torch_dtype=torch.float16)
        pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-2-1", scheduler=schedulersetting).to('cuda') #replaced with stablediffusion-v-1-4 git repo #"CompVis/stable-diffusion-v1-4",use_auth_token=True
        torch.manual_seed(seed)
    if userInputScheduler == "ddim":
        schedulersetting = DDIMScheduler.from_pretrained("./stable-diffusion-2-1", subfolder="scheduler", revision="fp16",torch_dtype=torch.float16)
        pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-2-1", scheduler=schedulersetting).to('cuda') #replaced with stablediffusion-v-1-4 git repo #"CompVis/stable-diffusion-v1-4",use_auth_token=True
        torch.manual_seed(seed)
    if userInputScheduler == "lmsd":
        schedulersetting = LMSDiscreteScheduler.from_pretrained("./stable-diffusion-2-1", subfolder="scheduler", revision="fp16",torch_dtype=torch.float16)
        pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-2-1", scheduler=schedulersetting).to('cuda') #replaced with stablediffusion-v-1-4 git repo #"CompVis/stable-diffusion-v1-4",use_auth_token=True
        torch.manual_seed(seed)
    if userInputScheduler == "pndm":
        schedulersetting = PNDMScheduler.from_pretrained("./stable-diffusion-2-1", subfolder="scheduler", revision="fp16",torch_dtype=torch.float16, steps_offset=1, clip_sample=False)
        pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-2-1", scheduler=schedulersetting).to('cuda') #replaced with stablediffusion-v-1-4 git repo #"CompVis/stable-diffusion-v1-4",use_auth_token=True
        torch.manual_seed(seed)
    
    userGuidanceScale = userInputGuidance

    prompt = promtInput
    image = pipe(prompt, height=inputheight, width=inputwidth, num_inference_steps=iterationsNum, guidance_scale=userGuidanceScale, eta=0.0).images[0]
    image.save(saveName)
