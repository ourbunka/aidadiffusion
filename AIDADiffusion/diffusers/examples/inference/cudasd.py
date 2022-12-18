import random
from typing import List, Optional, Union


import torch
import os
import sys
# Add vendor directory to module search path
parent_dir = os.path.abspath(os.path.dirname(__file__))
vendor_dir = os.path.join(parent_dir, 'Vendor')

sys.path.append(vendor_dir)
# import torch  #pip3 install torch torchvision torchaudio --extra-index-url https://download.pytorch.org/whl/cu117 and then copy to vendor directory


def cudasdfunc(promtInput, iterationsNum, saveName ):
    torch.backends.cudnn.benchmark = True
    torch.backends.cuda.matmul.allow_tf32 = True

    seed = random.randint(0,9999999999999)
    print(seed)
    print(torch.__version__)
    
    lms = LMSDiscreteScheduler(beta_start=0.00085, beta_end=0.012, beta_schedule="scaled_linear")
    pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-v1-4", scheduler=lms).to("cuda") #replaced with stablediffusion-v-1-4 git repo #"CompVis/stable-diffusion-v1-4",use_auth_token=True
    torch.manual_seed(seed)


    prompt = promtInput
    with torch.autocast("cuda"):
        image = pipe(prompt, height=768, width=512, num_inference_steps=iterationsNum, guidance_scale=15, eta=0.0)["sample"][0]
    image.save(saveName)