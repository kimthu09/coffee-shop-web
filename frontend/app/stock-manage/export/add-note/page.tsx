"use client";

import ExportIngredient from "@/components/stock-manage/export-ingredient";
import IngredientInsert from "@/components/stock-manage/ingredient-insert";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useForm } from "react-hook-form";

export type FormValues = {
  ingredients: {
    idIngre: string;
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
        <h1 className="font-medium text-xxl self-start">Thêm phiếu xuất</h1>
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
                  <Label htmlFor="lyDo">Lý do</Label>
                  <Input id="lyDo"></Input>
                </div>
              </CardContent>
            </Card>
            <Card>
              <CardContent className="p-6">
                <Label>Thông tin nguyên liệu</Label>
                <ExportIngredient form={form} />
              </CardContent>
            </Card>
          </div>
        </form>
      </div>
    </div>
  );
};

export default AddNote;
