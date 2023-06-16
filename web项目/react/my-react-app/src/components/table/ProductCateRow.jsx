export default function ({ cate }) {
  return (
    <>
      <tr>
        <td
          style={{ color: "red", fontSize: "19px", textAlign: "center" }}
          colSpan={2}
        >
          {cate}
        </td>
      </tr>
    </>
  );
}
