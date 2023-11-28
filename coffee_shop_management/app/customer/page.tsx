import { CustomerTable } from "@/components/customer/customer-table";
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
import Link from "next/link";
import React from "react";

const CustomerManage = () => {
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Khách hàng</h1>
        <div>
          <Dialog>
            <DialogTrigger asChild>
              <Button>Thêm khách hàng</Button>
            </DialogTrigger>
            <DialogContent className="xl:max-w-[720px] max-w-[472px] p-0">
              <DialogHeader>
                <DialogTitle className="p-6 pb-0">
                  Thêm danh mục sản phẩm
                </DialogTitle>
              </DialogHeader>
              <div className="border-y-[1px] p-6 flex flex-col gap-4">
                <div>
                  <Label htmlFor="name">Tên khách hàng</Label>
                  <Input id="name"></Input>
                </div>
                <div>
                  <Label htmlFor="phone">Điện thoại</Label>
                  <Input id="phone"></Input>
                </div>
                <div>
                  <Label htmlFor="email">Email</Label>
                  <Input id="email"></Input>
                </div>
              </div>
              <DialogFooter className="p-6 pt-0">
                <DialogClose asChild>
                  <div>
                    <Button type="submit">Thêm</Button>
                  </div>
                </DialogClose>
              </DialogFooter>
            </DialogContent>
          </Dialog>
        </div>
      </div>

      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <CustomerTable />
      </div>
    </div>
  );
};

export default CustomerManage;
