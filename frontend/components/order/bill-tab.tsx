import { FormValues } from "@/app/order/page";
import React from "react";
import {
  Control,
  FieldArrayWithId,
  UseFieldArrayRemove,
  UseFormRegister,
  UseFormSetValue,
  UseFormWatch,
  useWatch,
} from "react-hook-form";
import { Card, CardContent } from "../ui/card";
import { FiTrash2 } from "react-icons/fi";
import { HiPlus, HiMinus } from "react-icons/hi";
import { FaPen } from "react-icons/fa";
import { PiClipboardTextLight } from "react-icons/pi";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { toVND } from "@/lib/utils";
import { Label } from "../ui/label";

const AddUp = ({
  control,
  index,
}: {
  control: Control<FormValues>;
  index: number;
}) => {
  const formValues = useWatch({
    name: `invoiceDetails.${index}`,
    control,
  });
  const addUp = formValues.price * formValues.quantity;
  return <span className="text-sm font-bold">{toVND(addUp)}</span>;
};

export const Total = ({ control }: { control: Control<FormValues> }) => {
  const formValues = useWatch({
    name: "invoiceDetails",
    control,
  });
  const total = formValues.reduce(
    (acc, current) => acc + (current.price || 0) * (current.quantity || 0),
    0
  );
  const totalQuantity = formValues.reduce(
    (acc, current) => acc + 1 * (current.quantity || 0),
    0
  );
  return (
    <div className="flex gap-2 items-center">
      <span>Tổng tiền</span>
      <div className="border rounded-lg px-2 py-1">{totalQuantity}</div>
      <h1 className="text-sm">{toVND(total)}</h1>
    </div>
  );
};

const BillTab = ({
  fields,
  setValue,
  register,
  watch,
  control,
  remove,
  isSheet,
}: {
  fields: FieldArrayWithId<FormValues, "invoiceDetails", "id">[];
  setValue: UseFormSetValue<FormValues>;
  register: UseFormRegister<FormValues>;
  watch: UseFormWatch<FormValues>;
  control: Control<FormValues, any>;
  remove: UseFieldArrayRemove;
  isSheet?: boolean;
}) => {
  const invoices = watch("invoiceDetails");
  return (
    <Card className="sticky right-0 top-0 h-[86vh] overflow-hidden">
      <CardContent
        className={`flex flex-col p-0 overflow-hidden h-[86vh] ${
          isSheet ? "rounded-none" : ""
        }`}
      >
        <div className="flex flex-col bg-white  shadow-[0_2px_2px_-2px_rgba(0,0,0,0.2)]">
          <div className="p-4">
            <Input placeholder="Tìm kiếm khách hàng"></Input>
          </div>
        </div>
        <div className="flex flex-col gap-2  overflow-auto pt-4 flex-1">
          {fields.length < 1 ? (
            <div className="flex flex-col items-center gap-4 py-8 text-muted-foreground font-medium pt-[20%]">
              <PiClipboardTextLight className="h-24 w-24 text-muted-foreground/40" />
              <span>Chọn sản phẩm</span>
            </div>
          ) : null}
          {fields.map((item, index) => {
            return (
              <div
                key={item.id}
                className={`flex ${
                  index === fields.length - 1 ? "" : "border-b"
                }  xl:px-4 px-2 pb-2 group gap-2`}
              >
                <span className="text-sm leading-6">{index + 1}</span>
                <div className="flex flex-col flex-1">
                  {/* Name size price row */}
                  <div className="flex">
                    <div className="flex basis-[35%]">
                      <h1 className="text-base font-medium">{item.foodName}</h1>
                    </div>

                    <div className="flex flex-wrap basis-[65%] items-center justify-between xl:gap-3 gap-2">
                      {/* Quantity */}
                      <div className="flex gap-2 items-center">
                        <Button
                          className="p-[2px] bg-primary hover:bg-primary/90 rounded-full cursor-pointer text-white invisible  group-hover:visible h-5 w-5"
                          onClick={() => {
                            const quantity = +invoices.at(index)?.quantity!;
                            if (quantity === 1) {
                              //TODO: remove
                            } else {
                              setValue(
                                `invoiceDetails.${index}.quantity`,
                                quantity - 1
                              );
                            }
                          }}
                        >
                          <HiMinus />
                        </Button>
                        <Input
                          type="number"
                          className="px-1 w-10 text-center [&::-webkit-inner-spin-button]:appearance-none"
                          {...register(
                            `invoiceDetails.${index}.quantity` as const
                          )}
                        ></Input>

                        <Button
                          className="p-[2px] bg-primary hover:bg-primary/90 rounded-full cursor-pointer text-white invisible group-hover:visible h-5 w-5"
                          onClick={() => {
                            setValue(
                              `invoiceDetails.${index}.quantity`,
                              +invoices.at(index)?.quantity! + 1
                            );
                          }}
                        >
                          <HiPlus />
                        </Button>
                      </div>
                      <span className="font-semibold text-sm text-center">
                        M
                      </span>
                      <span className="text-sm text-right">
                        {toVND(item.price)}
                      </span>
                    </div>
                  </div>
                  {/* topping list */}
                  <div className="flex flex-col">
                    <span className="font-medium text-sm">Topping:</span>
                    <div className="flex justify-between">
                      <span className="text-muted-foreground text-sm">
                        Hat sen
                      </span>
                      <span className="text-muted-foreground text-sm">
                        5.000 d
                      </span>
                    </div>
                    <div className="flex justify-between">
                      <span className="text-muted-foreground text-sm">
                        Hat sen
                      </span>
                      <span className="text-muted-foreground text-sm">
                        5.000 d
                      </span>
                    </div>
                  </div>

                  <div className="flex">
                    <span className="text-sm font-semibold">Tổng</span>
                    <div className="text-right ml-auto">
                      <AddUp control={control} index={index} />
                    </div>
                  </div>
                  <div className="flex justify-between">
                    <div className="relative flex-1 items-center">
                      <span className="w-full">
                        <input
                          id={`note${index}`}
                          className="outline-none border-0 text-sm ml-5 w-4/5"
                          placeholder="Ghi chú..."
                        ></input>
                      </span>

                      <Label htmlFor={`note${index}`}>
                        <FaPen className="text-muted-foreground h-3 absolute top-2 cursor-pointer" />
                      </Label>
                    </div>
                    <Button
                      variant={"ghost"}
                      className="h-8 p-0 px-2 rounded-lg"
                      onClick={() => remove(index)}
                    >
                      <FiTrash2 className="opacity-50" />
                    </Button>
                  </div>
                </div>
              </div>
            );
          })}
        </div>
        <div className="flex justify-between items-center shadow-[0_-2px_2px_-2px_rgba(0,0,0,0.2)] bg-white h-20 px-4">
          <Button>Thanh toán</Button>
          {/* Total */}
          <div className="ml-auto">
            <Total control={control} />
          </div>
        </div>
      </CardContent>
    </Card>
  );
};

export default BillTab;
