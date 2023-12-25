import { CategoryTable } from "@/components/product-manage/category-table";
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
import React from "react";

const Category = () => {
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Danh mục</h1>
        <div>
          <Dialog>
            <DialogTrigger asChild>
              <Button>Thêm danh mục</Button>
            </DialogTrigger>
            <DialogContent className="xl:max-w-[720px] max-w-[472px] p-0">
              <DialogHeader>
                <DialogTitle className="p-6 pb-0">
                  Thêm danh mục sản phẩm
                </DialogTitle>
              </DialogHeader>
              <div className="border-y-[1px]">
                <div className="p-6">
                  <Label htmlFor="txtCate">Tên danh mục</Label>
                  <Input id="txtCate"></Input>
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
        <CategoryTable />
      </div>
    </div>
  );
};

export default Category;
