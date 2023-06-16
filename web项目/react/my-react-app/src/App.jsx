import { Component, useEffect, useMemo, useState } from "react";
import ProductTable from "./components/table/ProductTable";
import SearchBar from "./components/search/searchBar";
import "./App.css";
let data = [
  { category: "Fruits", price: "$1", stocked: true, name: "Apple" },
  { category: "Fruits", price: "$1", stocked: true, name: "Dragonfruit" },
  { category: "Fruits", price: "$2", stocked: false, name: "Passionfruit" },
  { category: "Vegetables", price: "$2", stocked: true, name: "Spinach" },
  { category: "Vegetables", price: "$4", stocked: false, name: "Pumpkin" },
  { category: "Vegetables", price: "$1", stocked: true, name: "Peas" },
];
function App() {
  // 输入框内容
  const [inputVale, setInputValue] = useState("");
  // 是否只显示有货
  const [isStocked, setIsStocked] = useState(false);

  // useEffect(() => console.log(1), [inputVale]);
  return (
    <>
      <SearchBar
        inputVale={inputVale}
        isStocked={isStocked}
        setInputValue={setInputValue}
        setIsStocked={setIsStocked}
      />
      <ProductTable
        products={data}
        inputVale={inputVale}
        isStocked={isStocked}
      />
    </>
  );
}

export default App;
