from cudasd import cudasdfunc
import sys


promtInput = str(sys.argv[1])
iterationsNum = int(sys.argv[2])
saveName = str(sys.argv[3] + ".jpg")
inputwidth = int(sys.argv[4])
inputheight = int(sys.argv[5])

print(promtInput, iterationsNum, saveName, inputwidth, inputheight)

cudasdfunc(promtInput, iterationsNum,saveName, inputwidth, inputheight)
