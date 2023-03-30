# AIDAdiffusion
Concept Art/Prototyping faster with ourbunka internal tool AIDA (AIDAdiffusion), "All-In-one" app for running Stable Diffusion on windows PC locally, with easy to use web-based UI!

Powered by Stable Diffusion.

<img src="https://github.com/ourbunka/aidadiffusion/blob/main/AIDA.PNG?raw=true">

# V2-1-2 
  -update to pytorch 2.0 cu118
  
  -~30% faster inference speed on ada lovelace gpus  


# Requirements

-nvidia gpu

-8GB of VRAM or more

-32GB of RAM or more 

-windows 10 or 11

-Python 3.10.7 installed

# Preview

[AIDA GUI github pages preview](https://ourbunka.github.io/aidadiffusion)

Watch [v0.1.2 video Demo here.](https://www.youtube.com/watch?v=1lm7o4PX-rI)


# ToS
BY USING AIDA(AIDADIFFUSION) APP, YOU ARE AGREEING WITH OURBUNKA MOTION AIDA(AIDADIFFUSION) TERMS OF SERVICES, LICENSE AND
[STABLE DIFFUSION LICENSE](https://huggingface.co/spaces/CompVis/stable-diffusion-license).

# Download
Download [here](https://drive.google.com/file/d/1yhdYxE57lHIB5-N-gQKv_Z9e_Yj2G4dw/view?usp=share_link)
also available on [itch.io](https://ourbunka.itch.io/aida)

# V2-1 Dependency installation
If you download Latest V2-1-Portable releases, ignore this. All python & stablediffusion dependencies in included in the release.

If you gitclone this repo, you need to copy these python libraries into aidadiffusion/diffusers/vendor directory
-diffusers
-torch ( pip3 install torch torchvision torchaudio --index-url https://download.pytorch.org/whl/cu118)


you also need to accept stablediffusion license on huggingface, download stable-diffusion-v2-1 repo. Then, copy to aidadiffusion/diffusers/


no huggingface token auth required after these step completed.


# Credits
We want to thanks authors, researchers and engineers of following repos for their research and implementations for making this app possible.

-@runwayml, @compvis, Robin Rombach, Andreas Blattmann, Domninik Lorenz, Patrick Esser, Björn Ommer, Stablility AI and Runway works on stable-diffusion https://github.com/CompVis/stable-diffusion & https://huggingface.co/runwayml/stable-diffusion-v1-5

-@harishanand95's diffusers https://github.com/harishanand95/diffusers

-@huggingface's diffusers https://github.com/huggingface/diffusers

-@chakra-ui's chakra-ui https://github.com/chakra-ui/chakra-ui
