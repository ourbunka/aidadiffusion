# AIDAdiffusion
Concept Art/Prototyping faster with ourbunka internal tool AIDA (AIDAdiffusion), "All-In-one" app for running Stable Diffusion on windows PC locally, with easy to use web-based UI🎉!

Powered by Stable Diffusion V1-4.

<img src="https://github.com/ourbunka/aidadiffusion/blob/main/AIDA.PNG?raw=true">


# Requirements

-modern gpu with DirectML support (Tested on AMD RX5700, newer nvidia & amd gpu should be supported too)

-10TFlops Fp32 or Higher for faster inferencing

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
Download AIDA(AIDAdiffusion) [Latest release here.](https://drive.google.com/file/d/1euC8uIzBq5D0mFpnncpb1CMoH1Mfyulz/view?usp=sharing) Also [available on Itch.io](https://ourbunka.itch.io/aida).

# V0.1.3 Dependency installation
If you download Latest V0.1.3-PrePack releases, ignore this. All python, stablediffusion & onnx dependencies in included in the release.

If you gitclone this repo, you need to copy these python libraries into aidadiffusion/diffusers/examples/inference/vendor directory
-diffusers
-ftfy
-numpy
-onnxruntime
-scipy
-torch
-tdm
-transformers

you also need to accept stablediffusion license on huggingface, download stable-diffusion-v-1-4 repo. Then, copy to aidadiffusion/diffusers/examples/inference/ . After that, convert the pretrained model to onnx runtime by [following this guide](https://gitgudblog.vercel.app/posts/stable-diffusion-amd-win10)

no huggingface token auth required after these step completed.

# Credits
We want to thanks authors, researchers and engineers of following repos for their research and implementations for making this app possible.

-@compvis, Robin Rombach, Andreas Blattmann, Domninik Lorenz, Patrick Esser, Björn Ommer, Stablility AI and Runway works on stable-diffusion https://github.com/CompVis/stable-diffusion

-@harishanand95's diffusers https://github.com/harishanand95/diffusers

-@huggingface's diffusers https://github.com/huggingface/diffusers

-@chakra-ui's chakra-ui https://github.com/chakra-ui/chakra-ui
