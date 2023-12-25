import { Button } from "@/components/ui/button";
import { Card, CardContent } from "@/components/ui/card";
import { Checkbox } from "@/components/ui/checkbox";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { roleFunctions } from "@/constants";
import React from "react";
import { LuCheck } from "react-icons/lu";

const AddRole = () => {
  return (
    <div className="col items-center">
      <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
        <div className="flex flex-row justify-between">
          <h1 className="font-medium text-xxl self-start">Thêm phân quyền</h1>
          <Button>
            <div className="flex flex-wrap gap-1 items-center">
              <LuCheck />
              Thêm
            </div>
          </Button>
        </div>
        <form className="flex flex-col gap-4">
          <Card>
            <CardContent className="p-6">
              <Label htmlFor="tenPhanQuyen">Tên phân quyền</Label>
              <Input id="tenPhanQuyen"></Input>
            </CardContent>
          </Card>
          <Card>
            <CardContent className="p-6">
              <div className="grid grid-cols-3 gap-y-6 gap-x-4">
                {roleFunctions.map((item) => {
                  return (
                    <div key={item.id} className="flex gap-2">
                      <Checkbox id={item.id}></Checkbox>
                      <Label>{item.name}</Label>
                    </div>
                  );
                })}
              </div>
            </CardContent>
          </Card>
        </form>
      </div>
    </div>
  );
};

export default AddRole;
