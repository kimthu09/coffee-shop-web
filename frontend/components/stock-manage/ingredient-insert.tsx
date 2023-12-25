import { useState } from "react";
import { Input } from "../ui/input";

import { Button } from "../ui/button";
import {
  Control,
  UseFormReturn,
  useFieldArray,
  useWatch,
} from "react-hook-form";

import { AiOutlineClose } from "react-icons/ai";
import { CiBoxes } from "react-icons/ci";
import { IoMdInformationCircleOutline } from "react-icons/io";
import { z } from "zod";
import { FormSchema } from "@/app/stock-manage/import/add-note/page";
import { AutoComplete } from "../autocomplete";
import { Ingredient } from "@/types";
import { toUnit, toVND } from "@/lib/utils";
import { toast } from "../ui/use-toast";
import getAllIngredient from "@/lib/getAllIngredient";
import Loading from "../loading";
const Total = ({
  control,
}: {
  control: Control<z.infer<typeof FormSchema>>;
}) => {
  const formValues = useWatch({
    name: "details",
    control,
  });
  const total = formValues.reduce(
    (acc, current) => acc + (current.price || 0) * (current.amountImport || 0),
    0
  );
  return <p>{toVND(total)}</p>;
};

const AddUp = ({
  control,
  index,
}: {
  control: Control<z.infer<typeof FormSchema>>;
  index: number;
}) => {
  const formValues = useWatch({
    name: `details.${index}`,
    control,
  });
  const addUp = formValues.price * formValues.amountImport;
  return <p>{toVND(addUp)}</p>;
};

const IngredientInsert = ({
  form,
}: {
  form: UseFormReturn<z.infer<typeof FormSchema>, any, undefined>;
}) => {
  const {
    register,
    handleSubmit,
    control,
    watch,
    getValues,
    reset,
    formState: { errors },
  } = form;
  const {
    fields: fieldsIngre,
    append: appendIngre,
    remove: removeIngre,
    replace,
  } = useFieldArray({
    control: control,
    name: "details",
  });
  const { data, isLoading, isError, mutate } = getAllIngredient();
  const [value, setValue] = useState<Ingredient>();
  const handleOnValueChange = (item: Ingredient) => {
    if (!fieldsIngre.find((ingre) => ingre.ingredientId === item.id)) {
      appendIngre({
        ingredientId: item.id,
        amountImport: 0,
        price: item.price,
        oldPrice: item.price,
        isReplacePrice: false,
      });
    }
  };
  if (isError) {
    return "Failed to fetch";
  } else if (isLoading || !data) {
    return <Loading />;
  } else {
    return (
      <div className="flex flex-col">
        <AutoComplete
          options={data.data}
          emptyMessage="Không có nguyên liệu khớp với từ khóa"
          placeholder="Tìm nguyên liệu"
          onValueChange={handleOnValueChange}
          value={value}
        />
        <div>
          <div className="grid grid-cols-4 lg:gap-3 gap-2 font-medium py-2 px-2 mt-2 rounded-t-md bg-[#ffe9db]">
            <h2 className="">Tên nguyên liệu</h2>
            <h2 className=" text-left">Đơn giá</h2>
            <h2 className=" text-left">Số lượng</h2>
            <h2 className=" text-right pr-12 ">Thành tiền</h2>
          </div>
          <div className="border border-t-0 py-2 rounded-b-md">
            {fieldsIngre.length < 1 ? (
              <div className="flex flex-col items-center gap-4 py-8 text-muted-foreground font-medium">
                <CiBoxes className="h-24 w-24 text-muted-foreground/40" />
                {errors.details?.root ? (
                  <span className="error___message">
                    {errors.details.root?.message}
                  </span>
                ) : (
                  "Chọn sản phẩm nhập kho"
                )}
              </div>
            ) : null}
            {fieldsIngre.map((ingre, index) => {
              const value = data.data.find(
                (item) => item.id === ingre.ingredientId
              );
              if (value) {
                return (
                  <div
                    key={ingre.id}
                    className="grid md:grid-cols-4 grid-cols-3 items-center p-2 lg:gap-3 gap-2"
                  >
                    <div className="flex">
                      <h2 className="font-medium">{value?.name}</h2>
                      <h2 className="ml-1 text-muted-foreground">
                        ({toUnit(value?.measureType)})
                      </h2>
                    </div>
                    <div className="relative p-1">
                      <Input
                        type="number"
                        defaultValue={ingre.price}
                        {...register(`details.${index}.price` as const)}
                        min={1}
                      ></Input>
                      <div className="absolute top-0 right-0 cursor-pointer group">
                        <IoMdInformationCircleOutline
                          className={`h-5 w-5 text-teal-700`}
                        />

                        <span
                          className="absolute bottom-5 right-3 w-fit whitespace-nowrap scale-0 transition-all rounded bg-teal-100 p-2 text-xs font-medium text-teal-800 group-hover:scale-100
                      group-active:scale-100"
                        >
                          Giá ban đầu: {toVND(ingre.oldPrice)}
                        </span>
                      </div>
                    </div>

                    <Input
                      type="number"
                      className="lg:w-full w-4/5"
                      defaultValue={ingre.amountImport}
                      {...register(`details.${index}.amountImport` as const)}
                      min={1}
                    ></Input>

                    <div className="text-right flex justify-end gap-2 items-center">
                      <AddUp control={control} index={index} />
                      <Button
                        type="button"
                        variant={"ghost"}
                        className={`px-3`}
                        onClick={() => {
                          removeIngre(index);
                        }}
                      >
                        <AiOutlineClose />
                      </Button>
                    </div>
                  </div>
                );
              } else {
                //TODO
              }
            })}
          </div>
        </div>

        <div className="flex justify-end pt-6 pr-14 font-medium ">
          <h2 className="w-1/4">Tổng cộng</h2>
          <div className="flex">
            <span>
              <Total control={control} />
            </span>
          </div>
        </div>
      </div>
    );
  }
};

export default IngredientInsert;
