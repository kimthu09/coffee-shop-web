"use client";

import IngredientInsert from "@/components/stock-manage/ingredient-insert";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useForm } from "react-hook-form";

export type FormValues = {
  ingredients: {
    idIngre: string;
    expirationDate: Date;
    quantity: number;
    price: number;
  }[];
};

const AddNote = () => {
  const form = useForm<FormValues>({
    defaultValues: {
      ingredients: [],
    },
  });
  const { register, handleSubmit, control, watch, getValues } = form;
  return (
    <div className="col items-center">
      <div className="col xl:w-4/5 w-full xl:px-0 md:px-6 px-0">
        <h1 className="font-medium text-xxl self-start">Thêm phiếu nhập</h1>
        <form>
          <div className="flex flex-col gap-4">
            <Card>
              <CardContent className="p-6 flex lg:flex-row flex-col gap-5">
                <div className="flex-1">
                  <Label htmlFor="idPhieu">Mã phiếu</Label>
                  <Input
                    id="idPhieu"
                    placeholder="Mã sinh tự động nếu để trống"
                  ></Input>
                </div>
                <div className="flex-1">
                  <Label htmlFor="idNcc">Nhà cung cấp</Label>
                  <Input id="idNcc"></Input>
                </div>
              </CardContent>
            </Card>
            <Card>
              <CardContent className="p-6">
                <Label>Thông tin nguyên liệu</Label>
                <IngredientInsert form={form} />
              </CardContent>
            </Card>
          </div>
        </form>
      </div>
    </div>
  );
};

export default AddNote;
