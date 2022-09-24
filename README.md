# AIDAdiffusion
Concept Art/Prototyping faster with ourbunka internal tool AIDA (AIDAdiffusion), "All-In-one" app for running Stable Diffusion on windows PC locally, with easy to use web-based UI🎉!

Powered by Stable Diffusion V1-4.

<img src="https://github.com/ourbunka/aidadiffusion/blob/main/AIDA.PNG?raw=true">

# V0.1.3 New Features
  -new Advance mode UI
  
  -more tweakable parameters such as custom seed, scheduler/sampler, guidance_scale etc.
  
  -new supported schedulers : LMS, DDIM,LMSD,PNDM
  
  -all dependencies included (If you download V0.1.3-Prepack)


# Requirements

-modern gpu with DirectML support (Tested on AMD RX5700, newer nvidia & amd gpu should be supported too)

-8GB of VRAM or more

-32GB of RAM or more 

-windows 10

-Python 3.10.7

# Preview

[AIDA GUI github pages preview](https://ourbunka.github.io/aidadiffusion)

Watch [vdeo Demo here.](https://www.youtube.com/watch?v=1lm7o4PX-rI)

*AIDA GUI required "AIDAdiffusion.exe" running in your local pc or cloud server to function properly

*currently github pages version seem to have issues connecting with locally hosted AIDAdiffusion.exe(github pages seem to automatically convert http to https?),
we recommend using locally hosted AIDA GUI [http://localhost:9998/files/build.](http://localhost:9998/files/build)

# ToS
BY USING AIDA(AIDADIFFUSION) APP, YOU ARE AGREEING WITH OURBUNKA MOTION AIDA(AIDADIFFUSION) TERMS OF SERVICES, LICENSE AND
[STABLE DIFFUSION LICENSE](https://huggingface.co/spaces/CompVis/stable-diffusion-license).

# Download
Download AIDA(AIDAdiffusion) [Latest release here.](https://drive.google.com/file/d/1lohvZilOsWE5DwvhkCu1ABvHdhVY0Y0u/view?usp=sharing) Also [available on Itch.io](https://ourbunka.itch.io/aida).

# V0.1.3 Dependency installation
If you download Latest V0.1.3-PrePack releases, ignore this. All python, stablediffusion & onnx dependencies in included in the release.

If you gitclone this repo, you need to copy these python libraries into aidadiffusion/diffusers/examples/inference/vendor directory
-diffusers
-ftfy
-numpy
-onnxruntime
-scipy
-torch
-tqdm
-transformers

you also need to accept stablediffusion license on huggingface, download stable-diffusion-v-1-4 repo. Then, copy to aidadiffusion/diffusers/examples/inference/ . After that, go to aidadiffusion/diffusers/examples/inference/ in terminal, run "python save_onnx.py" script to convert the pretrained model to onnx runtime.

no huggingface token auth required after these step completed.

# Credits
We want to thanks authors, researchers and engineers of following repos for their research and implementations for making this app possible.

-@compvis, Robin Rombach, Andreas Blattmann, Domninik Lorenz, Patrick Esser, Björn Ommer, Stablility AI and Runway works on stable-diffusion https://github.com/CompVis/stable-diffusion

-@harishanand95's diffusers https://github.com/harishanand95/diffusers

-@huggingface's diffusers https://github.com/huggingface/diffusers

-@chakra-ui's chakra-ui https://github.com/chakra-ui/chakra-ui
