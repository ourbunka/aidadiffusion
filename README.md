# Cuda port of AIDA (AIDAdiffusion)  AIDA-C
Under Developments

# Todo
-port to cuda to support Nvidia GPUs
-add float16 mode in ui

# AIDAdiffusion
Concept Art/Prototyping faster with ourbunka internal tool AIDA (AIDAdiffusion), "All-In-one" app for running Stable Diffusion on windows PC locally, with easy to use web-based UI🎉!

Powered by Stable Diffusion.

<img src="https://github.com/ourbunka/aidadiffusion/blob/main/AIDA.PNG?raw=true">

# V0.1.4 New Features
  -updated to runwayml/stable diffusion v-1-5
  -better and more realistic generated result
  -cuda and cpu only(can't get stablediffusion v1.5 to run on onnxruntime)
  -all dependencies included (If you download V0.1.4-Portable) *coming soon
  -smaller filesizes

# Requirements

-nvidia gpu
-cpu is supported, but it will be significantly slower(16s/it on a 6core12threads ryzen 3600 vs ~1sec/it on a 10Teraflops gpu)

-8GB of VRAM or more

-32GB of RAM or more 

-windows 10 or 11

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
coming soon

# V0.1.4 Dependency installation
If you download Latest V0.1.4-Portable releases, ignore this. All python, stablediffusion & onnx dependencies in included in the release.

If you gitclone this repo, you need to copy these python libraries into aidadiffusion/diffusers/examples/inference/vendor directory
-diffusers
-ftfy
-numpy
-scipy
-torch
-tdm
-transformers

you also need to accept stablediffusion license on huggingface, download stable-diffusion-v-1-5 repo. Then, copy to aidadiffusion/diffusers/examples/inference/ . After that, convert the pretrained model to onnx runtime by [following this guide](https://gitgudblog.vercel.app/posts/stable-diffusion-amd-win10)

no huggingface token auth required after these step completed.

# Credits
We want to thanks authors, researchers and engineers of following repos for their research and implementations for making this app possible.

-@runwayml, @compvis, Robin Rombach, Andreas Blattmann, Domninik Lorenz, Patrick Esser, Björn Ommer, Stablility AI and Runway works on stable-diffusion https://github.com/CompVis/stable-diffusion & https://huggingface.co/runwayml/stable-diffusion-v1-5

-@harishanand95's diffusers https://github.com/harishanand95/diffusers

-@huggingface's diffusers https://github.com/huggingface/diffusers

-@chakra-ui's chakra-ui https://github.com/chakra-ui/chakra-ui
