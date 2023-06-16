import { useState } from "react";

export default function ({ product }) {
  let [edite, setEdite] = useState(false);
  return (
    <>
      <tr style={{ textAlign: "center" }}>
        <td
          style={{
            color: product.stocked ? "" : "red",
            textAlign: "start",
          }}
        >
          {product.name}
        </td>
        <td onDoubleClick={() => setEdite(true)}>
          {!edite ? <span>{product.price}</span> : <input type="text" />}
        </td>
      </tr>
    </>
  );
}
