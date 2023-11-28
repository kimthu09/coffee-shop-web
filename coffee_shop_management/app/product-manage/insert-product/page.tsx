"use client";
import { Button } from "@/components/ui/button";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Card, CardContent } from "@/components/ui/card";

import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useEffect, useState } from "react";
import { useFieldArray, useForm } from "react-hook-form";
import { AiOutlineClose } from "react-icons/ai";
import { IoReloadSharp } from "react-icons/io5";
import IngredientTabs from "@/components/product-manage/ingredient-tabs";
import CategoryList from "@/components/category-list";
export type FormValues = {
  productPrice: {
    price: number;
    costPrice: number;
    name?: string;
  }[];
  ingredients: {
    idIngre: string;
    amount: number;
    priceId: number;
  }[];
};

const InsertProductPage = () => {
  const form = useForm<FormValues>({
    defaultValues: {
      productPrice: [{ price: 0, costPrice: 0 }],
      ingredients: [],
    },
  });
  const { register, handleSubmit, control, watch, getValues } = form;

  const { fields, append, remove } = useFieldArray({
    control: control,
    name: "productPrice",
  });

  const [openCategory, setOpenCategory] = useState(false);
  const [category, setCategory] = useState("");
  const [newCategory, setNewCategory] = useState("");

  const [sizeNames, setSizeNames] = useState<string[]>([]);

  // useEffect(() => {
  //   const subscription = watch((value) => {
  //     console.log(value);
  //   });
  //   return () => subscription.unsubscribe();
  // }, [watch]);

  const setTabs = () => {
    const currentSizeNames = getValues("productPrice").map(
      (item) => item.name || ""
    );
    setSizeNames(currentSizeNames);
    console.log(getValues("ingredients"));
  };

  const onFormSubmit = (data: FormValues) => {
    console.log(data);
  };

  return (
    <div className="col items-center">
      <div className="col w-full xl:px-16 sm:px-4 ">
        <h1 className="font-medium text-xxl self-start">Thêm sản phẩm</h1>
        <form onSubmit={handleSubmit(onFormSubmit)}>
          <div className="flex flex-col gap-4 xl:flex-row">
            <div className="xl:basis-3/5 flex flex-col gap-4">
              <Card>
                <CardContent className="p-6">
                  <div className="flex flex-col gap-5">
                    <Label htmlFor="prodName">Tên sản phẩm</Label>
                    <Input id="prodName"></Input>

                    <div>
                      {/* price list */}
                      {fields.map((field, index) => {
                        return (
                          <div key={field.id} className="flex ">
                            <div
                              className={`flex-1 mr-2 ${
                                fields.length > 1 ? "" : ""
                              }`}
                            >
                              <Label>Tên giá</Label>
                              <Input
                                type="text"
                                {...register(
                                  `productPrice.${index}.name` as const
                                )}
                              ></Input>
                            </div>
                            <div
                              className={`flex-1 ${
                                fields.length > 1 ? "mx-2" : "mr-2"
                              }  `}
                            >
                              <Label>Giá bán</Label>
                              <Input
                                type="number"
                                min={0}
                                max={500000}
                                required
                              ></Input>
                            </div>
                            <div className="flex-1 ml-2 mr-1">
                              <Label>Giá vốn</Label>
                              <Input
                                type="number"
                                min={0}
                                max={500000}
                                required
                              ></Input>
                            </div>

                            {index != 0 ? (
                              <Button
                                variant={"ghost"}
                                className={`self-end px-3 gap-0 `}
                                onClick={() => {
                                  if (index > 0) {
                                    remove(index);
                                  }
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
                                  if (index > 0) {
                                    remove(index);
                                  }
                                }}
                              >
                                <AiOutlineClose />
                              </Button>
                            )}
                          </div>
                        );
                      })}
                      <div>
                        <Label className="text-primary text-lg cursor-pointer">
                          +
                        </Label>
                        <Button
                          className="self-start p-2"
                          variant={"link"}
                          onClick={() => {
                            append({ price: 0, costPrice: 0 });
                          }}
                        >
                          Thêm giá
                        </Button>
                      </div>
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
                    <Button
                      variant={"ghost"}
                      className="rounded-full"
                      size="icon"
                      onClick={() => {
                        setTabs();
                      }}
                    >
                      <IoReloadSharp />
                    </Button>
                  </div>

                  {sizeNames.length > 0 ? (
                    <Tabs defaultValue={"0"}>
                      <TabsList className="w-full justify-start mb-2">
                        {getValues("productPrice").map((price, index) => (
                          <TabsTrigger key={index} value={index.toString()}>
                            {price.name || "Tên giá " + ++index}
                          </TabsTrigger>
                        ))}
                      </TabsList>

                      {getValues("productPrice").map((price, index) => (
                        <TabsContent key={index} value={index.toString()}>
                          <IngredientTabs priceId={index} form={form} />
                        </TabsContent>
                      ))}
                      <TabsContent value="account"></TabsContent>
                      <TabsContent value="password"></TabsContent>
                    </Tabs>
                  ) : null}
                </CardContent>
              </Card>
            </div>

            <Card className="xl:basis-2/5 xl:self-start">
              <CardContent className="p-6">
                <div className="flex gap-4 flex-col ">
                  {/* category list */}
                  <Label htmlFor="cateList">Danh mục</Label>
                  <CategoryList
                    category={category}
                    setCategory={setCategory}
                    canAdd={true}
                  />

                  <Label className="mt-4" htmlFor="picture">
                    Hình ảnh
                  </Label>
                  <Input id="picture" type="file" />
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
