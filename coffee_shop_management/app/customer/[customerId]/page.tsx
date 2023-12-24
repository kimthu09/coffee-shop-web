import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { customers } from "@/constants";
import React from "react";

const CustomerDetail = ({ params }: { params: { customerId: string } }) => {
  const customer = customers.find((item) => item.id === params.customerId);
  return (
    <div className="col items-center">
      <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
        <h1 className="xl:text-3xl text-2xl">Khách hàng: {customer?.id}</h1>
        <Card>
          <CardContent className="p-6 flex flex-col   gap-4">
            <div className="flex gap-4  flex-col">
              <div className="flex gap-4 lg:flex-row flex-col">
                <div className="basis-2/3">
                  <Label htmlFor="name">Tên khách hàng</Label>
                  <Input id="name" defaultValue={customer?.name}></Input>
                </div>
                <div className="basis-1/3">
                  <Label htmlFor="phone">Điện thoại</Label>
                  <Input id="phone" defaultValue={customer?.phone}></Input>
                </div>
              </div>

              <div>
                <Label htmlFor="email">Email</Label>
                <Input id="email" defaultValue={customer?.email}></Input>
              </div>

              <div className="flex flex-row gap-4">
                <div className="flex-1">
                  <Label htmlFor="soHoaDon">Hoá đơn</Label>
                  <Input id="soHoaDon" defaultValue={0}></Input>
                </div>
                <div className="flex-1">
                  <Label htmlFor="diem">Điểm tích luỹ</Label>
                  <Input id="diem" defaultValue={customer?.point}></Input>
                </div>
              </div>
            </div>
          </CardContent>
        </Card>

        <div className="my-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
          {/* TODO: invoice table */}
        </div>
      </div>
    </div>
  );
};

export default CustomerDetail;
