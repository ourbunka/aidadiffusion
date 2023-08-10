import random
from typing import List, Optional, Union


import os
import sys
# Add vendor directory to module search path
parent_dir = os.path.abspath(os.path.dirname(__file__))
vendor_dir = os.path.join(parent_dir, 'vendor')

sys.path.append(vendor_dir)
from diffusers import DiffusionPipeline, LMSDiscreteScheduler, DDIMScheduler, PNDMScheduler
import torch

def cudaadvfunc(promtInput, iterationsNum, saveName, userSeedInput, userInputScheduler, userInputGuidance, inputwidth, inputheight ):
    torch.backends.cuda.matmul.allow_tf32 = True
    seed = userSeedInput
    print(seed)

    #scheduler/samplers

    if userInputScheduler == "lms":
        schedulersetting = LMSDiscreteScheduler.from_pretrained("./stable-diffusion-2-1", subfolder="scheduler", revision="fp16",torch_dtype=torch.float16)
        pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-2-1", scheduler=schedulersetting).to('cuda') 
        torch.manual_seed(seed)
    if userInputScheduler == "ddim":
        schedulersetting = DDIMScheduler.from_pretrained("./stable-diffusion-2-1", subfolder="scheduler", revision="fp16",torch_dtype=torch.float16)
        pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-2-1", scheduler=schedulersetting).to('cuda') 
        torch.manual_seed(seed)
    if userInputScheduler == "lmsd":
        schedulersetting = LMSDiscreteScheduler.from_pretrained("./stable-diffusion-2-1", subfolder="scheduler", revision="fp16",torch_dtype=torch.float16)
        pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-2-1", scheduler=schedulersetting).to('cuda') 
        torch.manual_seed(seed)
    if userInputScheduler == "pndm":
        schedulersetting = PNDMScheduler.from_pretrained("./stable-diffusion-2-1", subfolder="scheduler", revision="fp16",torch_dtype=torch.float16, steps_offset=1, clip_sample=False)
        pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-2-1", scheduler=schedulersetting).to('cuda') 
        torch.manual_seed(seed)
    
    userGuidanceScale = userInputGuidance

    prompt = promtInput
    image = pipe(prompt, height=inputheight, width=inputwidth, num_inference_steps=iterationsNum, guidance_scale=userGuidanceScale, eta=0.0).images[0]
    image.save(saveName)
