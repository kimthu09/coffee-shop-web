import { Dialog, DialogTrigger } from "@radix-ui/react-dialog";
import React, { useState } from "react";
import { Button } from "../ui/button";
import { DialogClose, DialogContent, DialogTitle } from "../ui/dialog";
import { RadioGroup, RadioGroupItem } from "../ui/radio-group";
import { Label } from "../ui/label";
import { FaRegFileExcel } from "react-icons/fa6";
const ExportDialog = ({
  handleExport,
  setExportOption,
  isImport,
}: {
  handleExport: () => void;
  setExportOption: (value: string) => void;
  isImport: boolean;
}) => {
  const [open, setOpen] = useState(false);
  const onOpenChange = (open: boolean) => {
    if (open) {
      setExportOption("all");
    }
    setOpen(open);
  };
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogTrigger asChild>
        <Button variant={"outline"} className="p-2">
          <FaRegFileExcel className="mr-1 h-5 w-5 text-green-700" />
          <span>Xuất file</span>
        </Button>
      </DialogTrigger>
      <DialogContent className="p-0">
        <DialogTitle className="p-6 pb-0">
          Xuất danh sách {isImport ? "phiếu nhập" : "phiếu nợ"}
        </DialogTitle>
        <div className="flex flex-col border-y-[1px] p-6 gap-4">
          <Label>Giới hạn kết quả xuất</Label>
          <RadioGroup
            defaultValue="all"
            onValueChange={(e: string) => setExportOption(e)}
          >
            <div className="flex items-center space-x-2">
              <RadioGroupItem value="all" id="r1" />
              <Label htmlFor="r1" className="font-normal">
                Tất cả các phiếu
              </Label>
            </div>
            <div className="flex items-center space-x-2">
              <RadioGroupItem value="comfortable" id="r2" />
              <Label className="font-normal" htmlFor="r2">
                Các phiếu được chọn trong trang này
              </Label>
            </div>
          </RadioGroup>
        </div>

        <div className="ml-auto p-6 pt-0">
          <div className="flex gap-4">
            <Button
              type="button"
              variant={"outline"}
              onClick={() => setOpen(false)}
            >
              Thoát
            </Button>

            <Button
              type="button"
              onClick={() => {
                setOpen(false);
                handleExport();
              }}
            >
              Xuất
            </Button>
          </div>
        </div>
      </DialogContent>
    </Dialog>
  );
};

export default ExportDialog;
