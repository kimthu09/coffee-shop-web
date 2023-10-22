import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Sheet,
  SheetClose,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet";
import { FiUpload } from "react-icons/fi";

const ImportSheet = () => {
  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button className="hover:bg-orange-100 p-2" variant={"ghost"}>
          <div className="flex flex-wrap gap-1 items-center">
            <FiUpload />
            Nhập danh sách
          </div>
        </Button>
      </SheetTrigger>
      <SheetContent side={"top"} className="w-[480px] sm:w-[540px] m-auto">
        <SheetHeader>
          <SheetTitle>Nhập danh sách</SheetTitle>
          <SheetDescription>
            <div>
              <p>- Chuyển đổi file nhập dưới dạng .XLS trước khi tải dữ liệu</p>
              <p>
                <span>
                  - Tải file mẫu sản phẩm
                  <Button variant={"link"} className="px-1">
                    tại đây
                  </Button>
                </span>
              </p>
              <p>- File nhập có dung lượng tối đa là 3MB và 5000 bản ghi.</p>
            </div>
          </SheetDescription>
        </SheetHeader>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-4 items-center gap-4">
            <Input id="file" type="file" className="col-span-3" />
          </div>
        </div>
        <SheetFooter>
          <SheetClose asChild>
            <Button type="submit">Nhập file</Button>
          </SheetClose>
        </SheetFooter>
      </SheetContent>
    </Sheet>
  );
};

export default ImportSheet;
