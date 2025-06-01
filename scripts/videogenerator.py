# This script, unlike the audio generator, is mostly *inspired* by various online examples
# and resources. It works fine for its intended audio but it's probably sub-optimal otherwise.
# I wouldn’t use it beyond that without further refinement.
import numpy as np
import matplotlib.pyplot as plt
from scipy.io import wavfile
from matplotlib.animation import FuncAnimation, FFMpegWriter
from tqdm import tqdm
import subprocess

fps = 30
window_size = 2048

sample_rate, data = wavfile.read('song.wav')
data = data / np.max(np.abs(data))  # normalize audio

hop_size = int(sample_rate / fps)
freqs = np.fft.rfftfreq(window_size, d=1/sample_rate)
num_frames = (len(data) - window_size) // hop_size

# Precompute FFT magnitude frames for all time windows
fft_frames = []
for i in range(num_frames):
    start = i * hop_size
    windowed = data[start:start+window_size] * np.hanning(window_size)
    fft_mag = np.abs(np.fft.rfft(windowed))
    fft_frames.append(fft_mag)
fft_frames = np.array(fft_frames)
fft_frames /= np.max(fft_frames)  # normalize for plotting

# Set up plot appearance
fig, ax = plt.subplots(figsize=(12, 6))
fig.patch.set_facecolor('#0d0d0d')
ax.set_facecolor('#0d0d0d')
ax.axis('off')
ax.set_xlim(freqs[1], freqs[-1])
ax.set_ylim(0, 1.1)
ax.set_xscale('log')

line, = ax.plot(freqs, fft_frames[0], color='cyan', linewidth=1.5)

def update(frame):
    line.set_ydata(fft_frames[frame])
    return line,

anim = FuncAnimation(fig, update, frames=num_frames, interval=1000/fps, blit=True)

def progress_bar(anim, filename, writer):
    with tqdm(total=num_frames, desc="Saving video") as pbar:
        def progress_callback(i, n):
            pbar.update(1)
        anim.save(filename, writer=writer, progress_callback=progress_callback)

writer = FFMpegWriter(fps=fps, metadata=dict(artist='Me'), bitrate=1800)

video_filename = 'animated_fft.mp4'
output_filename = 'animated_fft_with_audio.mp4'

progress_bar(anim, video_filename, writer)
plt.close(fig)

# Merge video and original audio using ffmpeg
cmd = [
    'ffmpeg',
    '-y',  # overwrite output if exists
    '-i', video_filename,
    '-i', 'song.wav',
    '-c:v', 'copy',
    '-c:a', 'aac',
    '-b:a', '192k',
    '-shortest',
    output_filename
]

print("Merging audio and video...")
subprocess.run(cmd, check=True)
print(f"Done! Output saved to '{output_filename}'")
