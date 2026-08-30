import wrongUrl from "../assets/ya_oshibsya_tinkoff.mp3";
import correctUrl from "../assets/yra_pobeda.mp3";

let ctx = null;
const clips = {};

function clip(url) {
  if (typeof Audio === "undefined") return null;

  if (!clips[url]) {
    clips[url] = new Audio(url);
    clips[url].preload = "auto";
  }

  return clips[url];
}

// fallback - синтезированный сигнал, если mp3 не проигрался
function playClip(url, fallback) {
  const el = clip(url);
  if (!el) {
    fallback();
    return;
  }

  // сброс позиции - иначе быстрый второй ответ не переиграет звук
  el.currentTime = 0;
  el.play().catch(fallback);
}

function audio() {
  if (typeof window === "undefined") return null;

  const Ctx = window.AudioContext ?? window.webkitAudioContext;
  if (!Ctx) return null;

  if (!ctx) ctx = new Ctx();
  if (ctx.state === "suspended") ctx.resume();

  return ctx;
}

function tone(ac, freq, start, duration, gain = 0.16) {
  const osc = ac.createOscillator();
  const vol = ac.createGain();

  osc.type = "sine";
  osc.frequency.setValueAtTime(freq, ac.currentTime + start);

  vol.gain.setValueAtTime(0, ac.currentTime + start);
  vol.gain.linearRampToValueAtTime(gain, ac.currentTime + start + 0.012);
  vol.gain.exponentialRampToValueAtTime(
    0.0001,
    ac.currentTime + start + duration,
  );

  osc.connect(vol);
  vol.connect(ac.destination);

  osc.start(ac.currentTime + start);
  osc.stop(ac.currentTime + start + duration + 0.02);
}

function toneCorrect() {
  const ac = audio();
  if (!ac) return;

  tone(ac, 587.33, 0, 0.12);
  tone(ac, 880, 0.09, 0.22);
}

function toneWrong() {
  const ac = audio();
  if (!ac) return;

  tone(ac, 233.08, 0, 0.18, 0.13);
  tone(ac, 174.61, 0.13, 0.3, 0.13);
}

export function playCorrect() {
  playClip(correctUrl, toneCorrect);
}

export function playWrong() {
  playClip(wrongUrl, toneWrong);
}
