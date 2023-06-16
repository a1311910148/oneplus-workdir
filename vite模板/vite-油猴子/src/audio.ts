export default function () {
  const audio = new Audio("http://101.43.18.212:8883/王菲 - 匆匆那年.flac");
  document.body.appendChild(audio);
  audio.play();
}
