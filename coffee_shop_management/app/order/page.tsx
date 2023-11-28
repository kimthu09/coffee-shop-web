"use client";
import BillTab, { Total } from "@/components/order/bill-tab";
import ProductTab from "@/components/order/product-tab";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { FaChevronUp } from "react-icons/fa6";
import { useFieldArray, useForm } from "react-hook-form";
import { Sheet, SheetContent, SheetTrigger } from "@/components/ui/sheet";
export type FormValues = {
  invoiceDetails: {
    foodId: string;
    foodName: string;
    sizeName?: string;
    quantity: number;
    price: number;
  }[];
};
const OrderScreen = () => {
  const form = useForm<FormValues>({
    defaultValues: {
      invoiceDetails: [],
    },
  });
  const { register, control, setValue, watch } = form;

  const { fields, append, remove, update } = useFieldArray({
    control: control,
    name: "invoiceDetails",
  });
  return (
    <div className="flex gap-4 md:pb-0 pb-16">
      <div className="2xl:basis-3/5 xl:basis-1/2 md:basis-2/5  flex-1 ">
        <ProductTab append={append} />
      </div>
      <div className="2xl:basis-2/5 xl:basis-1/2 md:basis-3/5  md:block hidden">
        <BillTab
          fields={fields}
          setValue={setValue}
          register={register}
          watch={watch}
          control={control}
          remove={remove}
        />
      </div>
      <div className="fixed bottom-0 left-0 right-0">
        <Card className="md:hidden flex flex-col  h-16 bg-white rounded-none overflow-hidden">
          <div className="flex flex-1 justify-between items-center align-middle px-4">
            <Button>Thanh toán</Button>
            {/* Total */}
            <div className="ml-auto">
              <Total control={control} />
            </div>
          </div>

          <Sheet>
            <SheetTrigger asChild>
              <Button className="w-8 h-8 absolute p-0 rounded-full top-[-14px] left-[50%]">
                <FaChevronUp className="w-5 h-5" />
              </Button>
            </SheetTrigger>
            <SheetContent side={"bottom"} className="w-full p-0 bg-white pt-10">
              <BillTab
                fields={fields}
                setValue={setValue}
                register={register}
                watch={watch}
                control={control}
                remove={remove}
                isSheet
              />
            </SheetContent>
          </Sheet>
        </Card>
      </div>
    </div>
  );
};

export default OrderScreen;
