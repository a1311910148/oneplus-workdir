export default function ({ products }) {
  let stockedNum = products.filter((el) => el.stocked).length;
  return (
    <>
      <tr>
        <td>有货</td>
        <td style={{ textAlign: "center" }}>{stockedNum}</td>
      </tr>
    </>
  );
}
