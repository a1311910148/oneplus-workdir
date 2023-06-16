import { useState } from "react";
import "./App.css";
import BorderBox from "./components/cell/borderBox";
import CellTable from "./components/cell/cellTable";
import ContextMenu from "./components/cell/ContextMenu";
import dataRow from "./data";

function App() {
  // 单元格数据
  const [data, setData] = useState(dataRow);
  // 单元格输入处理函数
  function handleInput(id, value) {
    console.log(id, value);
    // 生成副本
    let newData = data.slice();
    // 修改对应id的数据
    newData[id].content = value;
    // 设置状态 更新组件
    setData(newData);
  }
  // 处理单元格选择
  function handleSelect(start, stop) {
    console.log(start, stop);
    // 生成副本
    let newData = data.slice();
    newData.forEach((el) => {
      el.select = false;
      if (
        el.x >= start.x &&
        el.x <= stop.x &&
        el.y >= start.y &&
        el.y <= stop.y
      ) {
        el.select = true;
      }
    });
    setData(newData);
  }
  return (
    <>
      <BorderBox>
        <CellTable
          data={data}
          handleInput={handleInput}
          onselect={handleSelect}
        />
      </BorderBox>
      <ContextMenu />
    </>
  );
}

export default App;
