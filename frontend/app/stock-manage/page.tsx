"use client";

import { StockTable } from "@/components/stock-manage/stock-table";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import UnitList from "@/components/unit-list";
import { useState } from "react";

const StockManage = () => {
  const [unit, setUnit] = useState("");
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Danh sách tồn kho</h1>
        <div>
          <Dialog>
            <DialogTrigger asChild>
              <Button>Tạo nguyên liệu</Button>
            </DialogTrigger>
            <DialogContent className="xl:max-w-[720px] max-w-[472px] p-0">
              <DialogHeader>
                <DialogTitle className="p-6 pb-0">Thêm nguyên liệu</DialogTitle>
              </DialogHeader>
              <div className="border-y-[1px]">
                <div className="p-6 xl:flex xl:gap-5">
                  <div className="flex-1 mb-4">
                    <Label htmlFor="txtNameIngre">Tên nguyên liệu</Label>
                    <Input id="txtNameIngre"></Input>
                  </div>

                  <div className="flex-1">
                    <Label htmlFor="txtNameIngre">Đơn vị</Label>
                    <UnitList unit={unit} setUnit={setUnit}></UnitList>
                  </div>
                </div>
              </div>
              <DialogFooter className="p-6 pt-0">
                <DialogClose asChild>
                  <div>
                    <Button type="submit">Lưu</Button>
                  </div>
                </DialogClose>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      </div>
      <div className="my-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <StockTable />
      </div>
    </div>
  );
};

export default StockManage;
