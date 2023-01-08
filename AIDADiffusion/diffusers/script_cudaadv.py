from cudaasd_adv import cudaadvfunc
import sys


promtInput = str(sys.argv[1])
iterationsNum = int(sys.argv[2])
saveName = str(sys.argv[3] + ".jpg")
userSeedInput =int(sys.argv[4])
userScheduler =str(sys.argv[5])
userGuidance = float(str(sys.argv[6]))
inputwidth = int(sys.argv[7])
inputheight = int(sys.argv[8])

print(promtInput, iterationsNum, saveName, userSeedInput, userScheduler, userGuidance, inputwidth, inputheight)

cudaadvfunc(promtInput, iterationsNum,saveName, userSeedInput, userScheduler, userGuidance, inputwidth, inputheight)