import { useState } from "react";

export default function ({ children }) {
  const [react, setReact] = useState({ left: 0, width: 0, top: 0, height: 0 });
  let arr = Array(50)
    .fill(1)
    .map((el, index) => index + 1);

  return (
    <>
      {/* x轴 */}
      <div className="top">
        {arr.map((el) => {
          return <div>{el}</div>;
        })}
      </div>
      <div style={{ display: "flex" }}>
        {/* y轴 */}
        <div className="left">
          {arr.map((el) => {
            return <div>{el}</div>;
          })}
        </div>
        {/* 单元格区域 */}
        <div>{children}</div>
      </div>
    </>
  );
}
