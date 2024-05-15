import sys
from scipy.io import wavfile
from pesq import pesq

refaudio=sys.argv[1]
degaudio=sys.argv[2]

rate, ref = wavfile.read(refaudio)
rate, deg = wavfile.read(degaudio)

print(pesq(rate, ref, deg, 'wb'))
