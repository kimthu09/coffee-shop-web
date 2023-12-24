"use client";
import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Card, CardContent } from "@/components/ui/card";

import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useEffect, useState } from "react";
import { useFieldArray, useForm } from "react-hook-form";
import { AiOutlineClose } from "react-icons/ai";
import IngredientTabs from "@/components/product-manage/ingredient-tabs";
import CategoryList from "@/components/category-list";
import { categories } from "@/constants";
export type FormValues = {
  productPrice: {
    price: number;
    costPrice: number;
    name?: string;
  }[];
  ingredients: {
    idIngre: string;
    amount: number;
    priceId: string;
  }[];
  categories: {
    idCate: string;
  }[];
};

const InsertProductPage = () => {
  const form = useForm<FormValues>({
    defaultValues: {
      productPrice: [{ price: 0, costPrice: 0 }],
      ingredients: [],
      categories: [],
    },
  });
  const { register, handleSubmit, control, watch } = form;

  const { fields, append, remove, update } = useFieldArray({
    control: control,
    name: "productPrice",
  });
  const {
    fields: fieldsCate,
    append: appendCate,
    remove: removeCate,
    update: updateCate,
  } = useFieldArray({
    control: control,
    name: "categories",
  });
  const prices = watch("productPrice");

  // const [category, setCategory] = useState("");

  // useEffect(() => {
  //   const subscription = watch((value) => {
  //     console.log(value);
  //   });
  //   return () => subscription.unsubscribe();
  // }, [watch]);

  const onFormSubmit = (data: FormValues) => {
    console.log(data);
  };

  // console.log(fieldsCate);
  return (
    <div className="col items-center">
      <div className="col w-full xl:px-16 sm:px-4 ">
        <h1 className="font-medium text-xxl self-start">Thêm sản phẩm</h1>
        <form onSubmit={handleSubmit(onFormSubmit)}>
          <Button>submit</Button>
          <div className="flex flex-col gap-4 xl:flex-row">
            <div className="xl:basis-3/5 flex flex-col gap-4">
              <Card>
                <CardContent className="p-6">
                  <div className="flex flex-col gap-4 2xl:flex-row 2xl:gap-2">
                    <div className="basis-2/5">
                      <Label htmlFor="masp">Mã sản phẩm</Label>
                      <Input
                        id="masp"
                        placeholder="Mã sinh tự động nếu để trống"
                      ></Input>
                    </div>
                    <div className="flex-1">
                      <Label htmlFor="prodName">Tên sản phẩm</Label>
                      <Input id="prodName"></Input>
                    </div>
                  </div>
                </CardContent>
              </Card>
              <Card>
                <CardContent className="p-6 pt-3">
                  <div className="flex flex-col gap-3">
                    {/* price list */}
                    {fields.map((field, index) => {
                      return (
                        <div key={field.id} className="flex gap-3">
                          <div className={`flex-1`}>
                            <Label>Tên giá</Label>
                            <Input
                              type="text"
                              {...register(
                                `productPrice.${index}.name` as const
                              )}
                              required
                            ></Input>
                          </div>
                          <div className={`flex-1  `}>
                            <Label>Giá bán (VND)</Label>
                            <Input
                              type="number"
                              min={0}
                              max={500000}
                              required
                              {...register(
                                `productPrice.${index}.price` as const
                              )}
                            ></Input>
                          </div>
                          <div className="flex-1">
                            <Label>Giá vốn (VND)</Label>
                            <Input
                              type="number"
                              min={0}
                              max={500000}
                              required
                              {...register(
                                `productPrice.${index}.costPrice` as const
                              )}
                            ></Input>
                          </div>

                          {fields.length > 1 ? (
                            <Button
                              variant={"ghost"}
                              className={`self-end px-3 gap-0 `}
                              onClick={() => {
                                remove(index);
                              }}
                            >
                              <AiOutlineClose />
                            </Button>
                          ) : (
                            <Button
                              variant={"ghost"}
                              className={`self-end px-3 gap-0 `}
                              disabled
                              onClick={() => {
                                remove(index);
                              }}
                            >
                              <AiOutlineClose />
                            </Button>
                          )}
                        </div>
                      );
                    })}
                    <div>
                      <Button
                        className="self-start p-2"
                        variant={"link"}
                        onClick={() => {
                          append({ price: 0, costPrice: 0 });
                        }}
                      >
                        <span className="font-bold">+</span> Thêm giá
                      </Button>
                    </div>
                  </div>
                </CardContent>
              </Card>
              {/* ingredient list */}
              <Card>
                <CardContent className="p-6 flex flex-col gap-4">
                  {/* ingredients list  */}
                  <div className="flex gap-2 items-center">
                    <Label>Nguyên liệu</Label>
                    {/* <Button
                      variant={"ghost"}
                      className="rounded-full"
                      size="icon"
                    >
                      <ChevronDown />
                    </Button> */}
                  </div>

                  {
                    <Tabs defaultValue={"0"}>
                      <TabsList className="w-full justify-start mb-2 h-fit flex-wrap">
                        {prices.map((price, index) => (
                          <TabsTrigger key={index} value={index.toString()}>
                            {price.name || "Tên giá "}
                          </TabsTrigger>
                        ))}
                      </TabsList>

                      {prices.map((price, index) => (
                        <TabsContent key={index} value={index.toString()}>
                          <IngredientTabs
                            priceId={fields.at(index)?.id!}
                            form={form}
                          />
                        </TabsContent>
                      ))}
                      <TabsContent value="account"></TabsContent>
                      <TabsContent value="password"></TabsContent>
                    </Tabs>
                  }
                </CardContent>
              </Card>
            </div>

            <Card className="xl:basis-2/5 xl:self-start">
              <CardContent className="p-6">
                <div className="flex gap-4 flex-col ">
                  {/* category list */}
                  <div>
                    <Label htmlFor="cateList">Danh mục</Label>
                    <CategoryList
                      checkedCategory={fieldsCate.map((cate) => cate.idCate)}
                      onCheckChanged={(idCate) => {
                        const selectedIndex = fieldsCate.findIndex(
                          (cate) => cate.idCate === idCate
                        );
                        if (selectedIndex > -1) {
                          removeCate(selectedIndex);
                        } else {
                          appendCate({ idCate: idCate });
                        }
                      }}
                      canAdd={true}
                    />
                    <div className="flex flex-wrap gap-2 mt-3">
                      {fieldsCate.map((cate, index) => (
                        <div
                          key={cate.id}
                          className="rounded-xl flex  px-3 py-1 h-fit outline-none text-sm text-primary  bg-orange-100 items-center gap-1 group"
                        >
                          {
                            categories.find((item) => item.id === cate.idCate)
                              ?.name
                          }
                          <div className="cursor-pointer w-4">
                            <AiOutlineClose className="group-hover:hidden" />
                            <AiOutlineClose
                              color="red"
                              fill="red"
                              className="text-primary group-hover:flex hidden h-4 w-4"
                              onClick={() => {
                                removeCate(index);
                              }}
                            />
                          </div>
                        </div>
                      ))}
                    </div>
                  </div>

                  <div>
                    <Label htmlFor="picture">Hình ảnh</Label>
                    <Input id="picture" type="file" />
                  </div>
                </div>
              </CardContent>
            </Card>
          </div>
        </form>
      </div>
    </div>
  );
};

export default InsertProductPage;
