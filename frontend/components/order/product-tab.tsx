"use client";
import { useEffect, useRef, useState } from "react";
import Image from "next/image";
import { Input } from "../ui/input";
import { categories, products } from "@/constants";
import { Product } from "@/types";
import { UseFieldArrayAppend, useFieldArray } from "react-hook-form";
import { FormValues } from "@/app/order/page";
import { removeAccents, toVND } from "@/lib/utils";
import { AspectRatio } from "@radix-ui/react-aspect-ratio";
import { useCallback } from "react";
const ProductTab = ({
  append,
}: {
  append: UseFieldArrayAppend<FormValues, "invoiceDetails">;
}) => {
  const [cateList, setCateList] = useState<
    {
      id: string;
      name: string;
      isSelected: boolean;
    }[]
  >();
  const [all, setAll] = useState(true);
  const [prodList, setProdList] = useState<Array<Product>>(products);
  const handleAllSelected = () => {
    if (!all) {
      setAll((prev) => !prev);
      setCateList(
        categories?.map((item: any) => {
          return { id: item.id, name: item.name, isSelected: false };
        })
      );
      setProdList(
        products?.filter((item: Product) => item.status === "active")
      );
    }
  };

  const handleCateSelected = (id: string) => {
    if (all) {
      setAll(false);
      setProdList(new Array());
    }
    const newCateList = cateList?.map((item: any) => {
      return {
        id: item.id,
        name: item.name,
        isSelected: item.id === id ? !item.isSelected : item.isSelected,
      };
    });
    setCateList(newCateList);

    if (newCateList?.every((item) => !item.isSelected)) {
      handleAllSelected();
    } else {
      const categorySet = new Set(
        newCateList
          ?.filter((item: any) => item.isSelected === true)
          .map((value) => value.id)
      );
      const newProdList = new Array<Product>();
      products.forEach((prod: Product) => {
        // for (let element of prod.idCate) {
        //   if (categorySet.has(element.id)) {
        //     if (book.quantity > 0) {
        //       newProdList.push(book);
        //       break;
        //     }
        //   }
        // }
      });
      setProdList(newProdList);
    }
  };

  const [inputValue, setInputValue] = useState<string>("");
  const [filteredList, setFilteredList] = useState<Array<Product>>(products);

  // Search Handler
  const searchHandler = useCallback(() => {
    const filteredData = products?.filter((prod) => {
      return removeAccents(prod.name)
        .toLowerCase()
        .includes(removeAccents(inputValue).toLowerCase().normalize());
    });
    setFilteredList(filteredData);
  }, [prodList, inputValue]);

  // EFFECT: Search Handler
  useEffect(() => {
    // Debounce search handler
    const timer = setTimeout(() => {
      searchHandler();
    }, 500);

    // Cleanup
    return () => {
      clearTimeout(timer);
    };
  }, [searchHandler]);

  const inputRef = useRef<HTMLInputElement>(null);

  useEffect(() => {
    document.addEventListener("keydown", detectKeyDown, true);
  }, []);
  const detectKeyDown = (e: any) => {
    if (e.key === "F2") {
      inputRef.current?.focus();
      setInputValue("");
    }
    return () => {
      document.removeEventListener("keydown", detectKeyDown);
    };
  };

  return (
    <div className="flex flex-col gap-6">
      <div className="flex items-end">
        <Input
          ref={inputRef}
          className=" bg-white rounded-xl"
          placeholder="Tìm kiếm sản phẩm"
          value={inputValue}
          onChange={(e) => {
            setInputValue(e.target.value);
          }}
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

            `}
            // onClick={() => handleCateSelected(index)}
          >
            {item.name}
          </div>
        ))}
      </div>
      <h1 className="text-lg">Sản phẩm</h1>

      {/* Product list */}
      <div className="grid 2xl:grid-cols-5 xl:grid-cols-4 lgr:grid-cols-3 md:grid-cols-2 sm:grid-cols-4 grid-cols-3 gap-4">
        {filteredList?.map((prod) => {
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
                  sizes="(max-width: 768px) 33vw, 20vw"
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
