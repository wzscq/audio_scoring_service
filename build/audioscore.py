import sys
import os
from scipy.io import wavfile
from pesq import pesq

refaudio=sys.argv[1]
refoutput = os.path.basename(refaudio)
refoutput = os.path.splitext(refoutput)[0]
refoutput = refoutput + "_refoutput.wav"
#转码
os.system("ffmpeg -i "+refaudio+" -ac 1 -ar 16000 -f wav  "+refoutput)

degaudio=sys.argv[2]
degoutput = os.path.basename(degaudio)
degoutput = os.path.splitext(refoutput)[0]
degoutput = refoutput + "_degoutput.wav"
#转码
os.system("ffmpeg -i "+degaudio+" -ac 1 -ar 16000 -f wav  "+degoutput)

rate, ref = wavfile.read(refoutput)
rate, deg = wavfile.read(degoutput)

print(pesq(rate, ref, deg, 'wb'))

#删除临时文件
os.remove(refoutput)
os.remove(degoutput)
