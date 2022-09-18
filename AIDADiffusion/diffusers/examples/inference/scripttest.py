from dml_onnxfunc import onnxfunc
import sys


promtInput = str(sys.argv[1])
iterationsNum = int(sys.argv[2])
saveName = str(sys.argv[3] + ".jpg")

print(promtInput, iterationsNum, saveName)

onnxfunc(promtInput, iterationsNum,saveName)