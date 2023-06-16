import ProductCateRow from "./ProductCateRow";
import ProductRow from "./ProductRow";
import CurrentBottom from "./CurrentBottom";

export default function ({ products, inputVale, isStocked }) {
  // 当前产品分类
  let category = null;
  if (inputVale) {
    products = products.filter((el) => el.name == inputVale);
  }
  if (isStocked) {
    products = products.filter((el) => el.stocked == true);
  }
  return (
    <>
      <table style={{ margin: "auto", width: "100%" }}>
        <thead>
          <tr>
            <th style={{ textAlign: "start", width: "80%" }}>name</th>
            <th>price</th>
          </tr>
        </thead>
        <tbody>
          {products.map((el) => {
            if (el.category != category) {
              category = el.category;
              return (
                <>
                  <ProductCateRow cate={el.category} />
                  <ProductRow product={el} />
                </>
              );
            } else {
              category = el.category;
              return <ProductRow product={el} />;
            }
          })}
          <CurrentBottom products={products} />
        </tbody>
      </table>
    </>
  );
}
