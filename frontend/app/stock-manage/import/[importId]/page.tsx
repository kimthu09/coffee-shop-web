"use client";
import ConfirmDialog from "@/components/confirm-dialog";
import Loading from "@/components/loading";
import { ExportImportNoteDetail } from "@/components/stock-manage/excel-import-detail";
import { ImportDetailTable } from "@/components/stock-manage/import-detail-table";
import { Button } from "@/components/ui/button";
import { toast } from "@/components/ui/use-toast";
import { getToken } from "@/lib/auth";
import updateStatus from "@/lib/import/changeStatus";
import getImportNoteDetail from "@/lib/import/getImportDetail";
import { toVND } from "@/lib/utils";
import withAuth from "@/lib/withAuth";
import { StatusNote } from "@/types";
import { useRouter } from "next/navigation";
import { BiBox } from "react-icons/bi";
import { FaRegFileExcel } from "react-icons/fa";
import { FiTrash2 } from "react-icons/fi";
import { LuPackageCheck, LuPhone } from "react-icons/lu";
const ImportDetail = ({ params }: { params: { importId: string } }) => {
  const token = getToken();
  const router = useRouter();
  if (!token) {
    router.push("/login");
  }
  const { data, isLoading, isError, mutate } = getImportNoteDetail({
    idNote: params.importId,
    token: token!,
  });
  const changeStatus = async (status: StatusNote) => {
    const response: Promise<any> = updateStatus({
      idNote: params.importId,
      status: status,
      token: token!,
    });
    const responseData = await response;
    console.log(responseData);
    if (responseData.hasOwnProperty("errorKey")) {
      toast({
        variant: "destructive",
        title: "Có lỗi",
        description: responseData.message,
      });
    } else {
      toast({
        variant: "success",
        title: "Thành công",
        description: "Chuyển trạng thái thành công",
      });
      mutate();
    }
  };
  if (isError) return <div>Failed to load</div>;
  if (isLoading) {
    return <Loading />;
  } else
    return (
      <div className="flex flex-col xl:mx-[20%] gap-6">
        <div className="shadow-sm bg-white flex flex-col gap-6 md:px-8 px-4 pb-6">
          <div className="flex justify-between gap-6 font-bold text-lg border-b flex-1 py-2 pt-6">
            <div className="flex gap-4">
              <span className="font-light">Mã phiếu nhập</span>
              <span>{data.id}</span>
            </div>
            <div className="flex gap-2">
              <Button
                variant={"outline"}
                className="p-2"
                onClick={() => {
                  ExportImportNoteDetail(data, data.details, "phieunhap.xlsx");
                }}
              >
                <FaRegFileExcel className="mr-1 h-5 w-5 text-green-700" />
                <span>Xuất file</span>
              </Button>
              <div
                className={`${
                  data.status === StatusNote.Inprogress ? "block" : "hidden"
                }`}
              >
                <ConfirmDialog
                  title={"Xác nhận hoàn thành phiếu nhập ?"}
                  description="Trạng thái sẽ không được thay đổi khi đã hoàn thành."
                  handleYes={() => changeStatus(StatusNote.Done)}
                >
                  <Button className={`p-2  bg-teal-600 hover:bg-teal-600/90`}>
                    <LuPackageCheck className="mr-1 h-6 w-6" />
                    <span>Hoàn thành</span>
                  </Button>
                </ConfirmDialog>
              </div>
              <ConfirmDialog
                title="Xác nhận hủy phiếu nhập ?"
                description="Trạng thái sẽ không được thay đổi khi đã hủy."
                handleYes={() => changeStatus(StatusNote.Cancel)}
              >
                <div
                  className={`${
                    data.status === StatusNote.Inprogress ? "block" : "hidden"
                  }`}
                >
                  <Button className="p-2 bg-red-500 hover:bg-red-500/90">
                    <FiTrash2 className="mr-1 h-5 w-5" />
                    <span>Hủy</span>
                  </Button>
                </div>
              </ConfirmDialog>
            </div>
          </div>
          <div className="grid grid-cols-2 text-sm">
            <div className="flex flex-col gap-4 w-fit">
              <div className="flex font-medium">
                <span className="w-16">Tạo</span>
                <div className="flex flex-col">
                  <span>{new Date(data.createdAt).toLocaleDateString()}</span>
                  <span className="font-light">{data.createdBy.name}</span>
                </div>
              </div>
              <div className="flex font-medium w-fit">
                <span className="w-16">Đóng</span>
                <div className="flex flex-col">
                  {data.closedAt ? (
                    <>
                      <span>12/10/2023</span>
                      <span className="font-light">Nguyễn Thị Huệ</span>
                    </>
                  ) : (
                    "Chưa đóng phiếu"
                  )}
                </div>
              </div>
            </div>
            <div className="flex flex-col items-end gap-4">
              <div className="w-fit">
                <div className="flex flex-col gap-2 font-medium">
                  <div className="flex items-center gap-1">
                    <BiBox className="h-4 w-4" />
                    {data.supplier.name}
                  </div>
                  <div className="flex items-center gap-2">
                    <LuPhone className="h-4 w-4" />
                    {data.supplier.phone}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="shadow-sm bg-white flex flex-col gap-6 py-6 md:px-6 px-4">
          <ImportDetailTable {...data.details} />
          <div className="flex justify-end space-x-2 pb-4 font-semibold">
            <span>Tổng tiền: </span>
            <span>{toVND(data.totalPrice)}</span>
          </div>
        </div>
      </div>
    );
};

export default withAuth(ImportDetail);
