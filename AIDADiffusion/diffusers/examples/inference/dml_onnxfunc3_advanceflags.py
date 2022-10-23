import random
import inspect
import warnings
from typing import List, Optional, Union



import os
import sys
# Add vendor directory to module search path
parent_dir = os.path.abspath(os.path.dirname(__file__))
vendor_dir = os.path.join(parent_dir, 'Vendor')

sys.path.append(vendor_dir)

import onnxruntime as ort
import torch
import numpy as np
from tqdm.auto import tqdm
from transformers import CLIPFeatureExtractor, CLIPTextModel, CLIPTokenizer
from diffusers.models import AutoencoderKL, UNet2DConditionModel
from diffusers.pipeline_utils import DiffusionPipeline
from diffusers.schedulers import DDIMScheduler, LMSDiscreteScheduler, PNDMScheduler, DDPMScheduler
# from diffusers import StableDiffusionSafetyChecker

def onnxfunc(promtInput, iterationsNum, saveName, userSeedInput, userInputScheduler, userInputGuidance ):
    seed = userSeedInput
    print(seed)

    device = torch.device("cuda") if torch.cuda.is_available() else torch.device("cpu")
    

    #scheduler/samplers
    

    ddim = DDIMScheduler(beta_start=0.00085, beta_end=0.012, beta_schedule="scaled_linear", tensor_format='numpy arrays')
    lmsd = LMSDiscreteScheduler(beta_start=0.00085, beta_end=0.012, beta_schedule="scaled_linear")
    pndm = PNDMScheduler(beta_start=0.00085, beta_end=0.012, beta_schedule="scaled_linear")
    ddpm = DDPMScheduler(beta_start=0.00085, beta_end=0.012, beta_schedule="scaled_linear", tensor_format='numpy arrays')
    lms = LMSDiscreteScheduler(beta_start=0.00085, beta_end=0.012, beta_schedule="scaled_linear")
    
    if userInputScheduler == "lms":
        pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-v1-5", revision="fp16",  scheduler=lms).to(device) #replaced with stablediffusion-v-1-4 git repo #"CompVis/stable-diffusion-v1-4",use_auth_token=True
        torch.manual_seed(seed)
    if userInputScheduler == "ddim":
        pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-v1-5", scheduler=ddim).to(device) #replaced with stablediffusion-v-1-4 git repo #"CompVis/stable-diffusion-v1-4",use_auth_token=True
        torch.manual_seed(seed)
    if userInputScheduler == "lmsd":
        pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-v1-5", scheduler=lmsd).to(device) #replaced with stablediffusion-v-1-4 git repo #"CompVis/stable-diffusion-v1-4",use_auth_token=True
        torch.manual_seed(seed)
    if userInputScheduler == "pndm":
        pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-v1-5", scheduler=pndm).to(device) #replaced with stablediffusion-v-1-4 git repo #"CompVis/stable-diffusion-v1-4",use_auth_token=True
        torch.manual_seed(seed)
    if userInputScheduler == "ddpm":
        pipe = DiffusionPipeline.from_pretrained("./stable-diffusion-v1-5", scheduler=ddpm).to(device) #replaced with stablediffusion-v-1-4 git repo #"CompVis/stable-diffusion-v1-4",use_auth_token=True
        torch.manual_seed(seed)
    
    userGuidanceScale = userInputGuidance

    prompt = promtInput
    image = pipe(prompt, height=768, width=768, num_inference_steps=iterationsNum, guidance_scale=userGuidanceScale, eta=0.0)["sample"][0]
    image.save(saveName)

# Works on AMD Windows platform
# Image width and height is set to 512x512
# If you need images of other sizes (size must be divisible by 8), make sure to save the model with that size in save_onnx.py
# For example, if you need height=512 and width=768, change create_onnx.py with height=512 and width=768 and run the prompt below with height=512 and width=768
# Default values are height=512, width=512, num_inference_steps=50, guidance_scale=7.5, eta=0.0, execution_provider="DmlExecutionProvider"

# prompt = "a photo of an astronaut riding a horse on mars"
# image = pipe(prompt, height=512, width=768, num_inference_steps=50, guidance_scale=7.5, eta=0.0, execution_provider="DmlExecutionProvider")["sample"][0]
# image.save("Dml_1.png")
