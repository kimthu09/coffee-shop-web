"use client";

import Link from "next/link";
import Image from "next/image";
import { Button } from "./ui/button";
import { LuLogOut } from "react-icons/lu";
import { useAuth } from "@/hooks/auth-context";
import ConfirmDialog from "./confirm-dialog";
const Header = () => {
  const { logout, user } = useAuth();
  return (
    <div className=" flex z-10 bg-white w-[100%] border-b border-gray-200">
      <div className="flex flex-1 h-[47px] items-center justify-between px-4">
        <div className="flex items-center space-x-4">
          <Link
            href="/"
            className="flex flex-row space-x-3 items-center justify-center md:hidden"
          >
            <div
              className={`flex align-middle justify-center items-center p-[4px] h-[40px] w-[40px]  rounded-xl bg-orange-100 `}
            >
              <Image
                src="/android-chrome-192x192.png"
                alt="logo"
                className="sidebar__logo"
                width={36}
                height={36}
              ></Image>
            </div>
            <span className="font-semibold text-lg flex ">Coffee Shop</span>
          </Link>
        </div>
        {user ? (
          <div className="flex px-6">
            <ConfirmDialog
              title={"Xác nhận hoàn thành phiếu nhập ?"}
              description="Trạng thái sẽ không được thay đổi khi đã hoàn thành."
              handleYes={() => logout()}
            >
              <Button variant={"link"}>
                <div className="flex gap-2 text-primary">
                  Đăng xuất
                  <LuLogOut className="h-5 w-5 " />
                </div>
              </Button>
            </ConfirmDialog>
          </div>
        ) : null}
      </div>
    </div>
  );
};

export default Header;
