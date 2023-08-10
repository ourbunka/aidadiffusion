import random
from typing import List, Optional, Union


import os
import sys


parent_dir = os.path.abspath(os.path.dirname(__file__))
vendor_dir = os.path.join(parent_dir, 'vendor')

sys.path.append(vendor_dir)
from diffusers import DiffusionPipeline
import torch




def cudasdfunc(promtInput, iterationsNum, saveName, inputwidth, inputheight ):
    #torch.backends.cudnn.benchmark = True
    torch.backends.cuda.matmul.allow_tf32 = True

    seed = random.randint(0,9999999999999)
    print(seed)
    print(torch.__version__)
   
    torch.manual_seed(seed)


    prompt = promtInput
    pipeline = DiffusionPipeline.from_pretrained("./stable-diffusion-2-1", torch_dtype=torch.float16)
    pipeline.to("cuda")
    image = pipeline(prompt, height=inputheight, width=inputwidth, num_inference_steps=iterationsNum, guidance_scale=7.5, eta=0.0).images[0]
    image.save(saveName)


