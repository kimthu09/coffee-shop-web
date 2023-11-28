"use client";
import { useState } from "react";
import Image from "next/image";
import { Input } from "../ui/input";
import { categories, products } from "@/constants";
import { Product } from "@/types";
import { UseFieldArrayAppend, useFieldArray } from "react-hook-form";
import { FormValues } from "@/app/order/page";
import { toVND } from "@/lib/utils";
import { AspectRatio } from "@radix-ui/react-aspect-ratio";
const ProductTab = ({
  append,
}: {
  append: UseFieldArrayAppend<FormValues, "invoiceDetails">;
}) => {
  const [cateList, setCateList] = useState<Array<Boolean>>(
    new Array(categories.length).fill(false)
  );
  const [all, setAll] = useState(true);
  const [prodList, setProdList] = useState<Array<Product>>(products);
  const handleAllSelected = () => {
    if (!all) {
      setAll((prev) => !prev);
      setCateList(new Array(categories.length).fill(false));
      setProdList(products);
    }
  };

  const handleCateSelected = (index: number) => {
    if (all) {
      setAll(false);
      setProdList(new Array());
    }
    const newCateList = cateList.map((item, idx) =>
      idx === index ? !item : item
    );
    setCateList(newCateList);

    if (newCateList.every((item) => !item)) {
      handleAllSelected();
    } else {
      const newProdList = products.filter((prod) =>
        categories
          .filter((item, idx) => newCateList[idx] === true)
          .find((cate) => cate.id === prod.idCate)
      );

      setProdList(newProdList);
    }
  };

  return (
    <div className="flex flex-col gap-6">
      <div className="flex items-end">
        <Input
          className=" bg-white rounded-xl"
          placeholder="Tìm kiếm sản phẩm"
        ></Input>
      </div>

      {/* Category list */}
      <div className="flex flex-wrap gap-2">
        <div
          className={` rounded-xl flex self-start px-3 py-1 border-gray-200 border text-sm  cursor-pointer ${
            all
              ? "bg-orange-50 border-primary text-brown font-medium"
              : "bg-white text-muted-foreground"
          }`}
          onClick={handleAllSelected}
        >
          Tất cả
        </div>
        {categories.map((item, index) => (
          <div
            key={item.id}
            className={`rounded-xl flex self-start px-3 py-1 border-gray-200 border text-sm  cursor-pointer
            ${
              cateList[index]
                ? "bg-orange-50 border-primary text-brown font-medium"
                : "bg-white text-muted-foreground"
            }`}
            onClick={() => handleCateSelected(index)}
          >
            {item.name}
          </div>
        ))}
      </div>
      <h1 className="text-lg">Sản phẩm</h1>

      {/* Product list */}
      <div className="grid 2xl:grid-cols-5 xl:grid-cols-4 lgr:grid-cols-3 md:grid-cols-2 sm:grid-cols-4 grid-cols-3 gap-4">
        {prodList.map((prod) => {
          return (
            <div
              key={prod.id}
              className="bg-white shadow-sm rounded-xl  overflow-hidden cursor-pointer hover:shadow-md"
              onClick={() => {
                append({
                  foodId: prod.id,
                  foodName: prod.name,
                  quantity: 1,
                  price: prod.price,
                });
              }}
            >
              <AspectRatio ratio={1 / 1}>
                <Image
                  className=" object-cover"
                  src={prod.image!}
                  alt="image"
                  fill
                ></Image>
              </AspectRatio>
              <div className="px-1">
                <h1 className="text-base font-medium text-center">
                  {prod.name}
                </h1>
                <h1 className="text-base font-semibold text-primary text-center pb-1">
                  {toVND(prod.price)}
                </h1>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};

export default ProductTab;
