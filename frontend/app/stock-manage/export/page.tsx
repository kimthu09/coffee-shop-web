import { ExportTable } from "@/components/stock-manage/export-table";
import Link from "next/link";
import React from "react";

const ImportStock = () => {
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Danh sách phiếu xuất kho</h1>
        <Link
          href="/stock-manage/export/add-note"
          className="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-9 px-4 py-2"
        >
          Thêm mới phiếu xuất
        </Link>
      </div>

      <div className="my-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <ExportTable />
      </div>
    </div>
  );
};

export default ImportStock;
