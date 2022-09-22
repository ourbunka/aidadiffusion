from dml_onnxfunc3_advanceflags import onnxfunc
import sys


promtInput = str(sys.argv[1])
iterationsNum = int(sys.argv[2])
saveName = str(sys.argv[3] + ".jpg")
userSeedInput =int(sys.argv[4])
userScheduler =str(sys.argv[5])
userGuidance = float(str(sys.argv[6]))

print(promtInput, iterationsNum, saveName, userSeedInput, userScheduler, userGuidance)

onnxfunc(promtInput, iterationsNum,saveName, userSeedInput, userScheduler, userGuidance)