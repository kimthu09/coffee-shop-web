import { StaffTable } from "@/components/staff/staff-table";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import React from "react";
import { FaPlus } from "react-icons/fa6";

const StaffManage = () => {
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Danh sách nhân viên</h1>
        <div>
          <Link
            href="staff/add"
            className="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-9 px-2 py-2"
          >
            <div className="flex flex-wrap gap-1 items-center">
              <FaPlus />
              Thêm nhân viên
            </div>
          </Link>
        </div>
      </div>
      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <StaffTable />
      </div>
    </div>
  );
};

export default StaffManage;
