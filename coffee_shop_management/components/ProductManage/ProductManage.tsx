import { Button } from "../ui/button";
import { FiDownload } from "react-icons/fi";
import { BiBox } from "react-icons/bi";
import { products } from "@/types";
import { ProductTable } from "./table";
import { Input } from "../ui/input";
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "../ui/dialog";
import ImportSheet from "./import-sheet";
const ProductManage = () => {
  const data = products;
  return (
    <div className="m-10">
      <div className="col">
        <div className="flex flex-row justify-between items-center">
          <h1 className="text-3xl font-bold tracking-tight lg:text-4xl">
            Sản phẩm
          </h1>
          <div>
            <Button>Thêm sản phẩm</Button>
          </div>
        </div>
        <div className="flex flex-row flex-wrap gap-2">
          <Button className="hover:bg-orange-100 p-2" variant={"ghost"}>
            <div className="flex flex-wrap gap-1 items-center">
              <FiDownload />
              Xuất file
            </div>
          </Button>

          <ImportSheet></ImportSheet>

          <Button className="hover:bg-orange-100 p-2" variant={"ghost"}>
            <div className="flex flex-wrap gap-1 items-center">
              <BiBox />
              Loại sản phẩm
            </div>
          </Button>
        </div>
      </div>

      <div className="mx-auto my-5 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
        <ProductTable></ProductTable>
      </div>
    </div>
  );
};

export default ProductManage;
