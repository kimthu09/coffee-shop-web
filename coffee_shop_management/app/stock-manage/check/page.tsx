import Link from "next/link";
import React from "react";

const CheckInventory = () => {
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-start">
        <h1>Tất cả phiếu kiểm kho</h1>
        <Link
          href="/stock-manage/import/add-note"
          className="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-9 px-4 py-2"
        >
          Thêm phiếu kiểm kho
        </Link>
      </div>

      <div className="my-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        {/* <ImportTable /> */}
      </div>
    </div>
  );
};

export default CheckInventory;
