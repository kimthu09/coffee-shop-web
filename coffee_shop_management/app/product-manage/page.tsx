import ImportSheet from "@/components/product-manage/import-sheet";
import { ProductTable } from "@/components/product-manage/table";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { FiDownload } from "react-icons/fi";

export default function ProductManage() {
  return (
    <div className="col">
      <div className="flex flex-row justify-between items-center">
        <h1>Sản phẩm</h1>
        <div>
          <Link
            href="/product-manage/insert-product"
            className="inline-flex items-center justify-center whitespace-nowrap rounded-md text-sm font-medium transition-colors focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:pointer-events-none disabled:opacity-50 bg-primary text-primary-foreground shadow hover:bg-primary/90 h-9 px-4 py-2"
          >
            Thêm sản phẩm
          </Link>
        </div>
      </div>
      <div className="flex flex-row flex-wrap gap-2">
        <Button className="hover:bg-orange-100 p-2" variant={"ghost"}>
          <div className="flex flex-wrap gap-1 items-center">
            <FiDownload />
            Xuất danh sách
          </div>
        </Button>

        <ImportSheet></ImportSheet>
      </div>

      <div className="mb-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <ProductTable></ProductTable>
      </div>
    </div>
  );
}
