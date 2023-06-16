export default function ({}) {
  window.addEventListener("contextmenu", (e) => {
    e.preventDefault();
    console.log(e);
  });
  return (
    <div className="ContextMenu">
      <ul>
        <li>复制</li>
        <li>粘贴</li>
        <li>颜色</li>
      </ul>
    </div>
  );
}
