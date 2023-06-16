export default function ({
  inputVale,
  isStocked,
  setInputValue,
  setIsStocked,
}) {
  return (
    <>
      <div>
        <input
          type="text"
          value={inputVale}
          onInput={(e) => setInputValue(e.target.value)}
        />
      </div>
      <div>
        <input
          type="checkbox"
          checked={isStocked}
          onChange={(e) => setIsStocked(e.target.checked)}
        />
        仅展示有货商品
      </div>
    </>
  );
}
