"use client";

import Image from "next/image";
import Link from "next/link";
import { RxHamburgerMenu } from "react-icons/rx";
import { createContext, useContext, useState } from "react";

import { usePathname } from "next/navigation";
import { ChevronDown } from "lucide-react";
import { sidebarItems } from "@/constants";
import { SidebarItem } from "@/types";
import { useAuth } from "@/hooks/auth-context";

const initialValue = {
  isCollapsed: false,
  toggleCollapse: () => {},
};
type SidebarType = {
  isCollapsed: boolean;
  toggleCollapse: () => void;
};
const SidebarContext = createContext<SidebarType | undefined>(undefined);

export const SidebarProvider = ({
  children,
}: {
  children: React.ReactNode;
}) => {
  const [isCollapsed, setCollapse] = useState<boolean>(false);
  const toggleCollapse = () => {
    setCollapse((prevState) => !prevState);
  };
  return (
    <SidebarContext.Provider
      value={{
        isCollapsed,
        toggleCollapse,
      }}
    >
      {children}
    </SidebarContext.Provider>
  );
};

export function SidebarNav() {
  const context = useContext(SidebarContext);
  const { user } = useAuth();
  if (!context) {
    throw new Error(
      "`SidebarContext` have to be used inside `CurtainContextProvider`"
    );
  }
  if (!user) {
    return null;
  } else
    return (
      <div className="md:flex hidden z-20">
        <aside
          className={`bg-white p-1 h-screen transition-all shadow-md overflow-auto ${
            context.isCollapsed ? "w-[3.8rem]" : "w-64"
          }`}
        >
          <nav>
            <div className={`flex items-center my-4 h-[64px]`}>
              <Link href="/">
                <div
                  className={`flex align-middle justify-center items-center gap-4 h-[64px] w-[64px]  rounded-xl bg-orange-100 ${
                    context.isCollapsed ? "hidden" : "flex"
                  }`}
                >
                  <Image
                    src="/android-chrome-192x192.png"
                    alt="logo"
                    className="sidebar__logo"
                    width={80}
                    height={80}
                  ></Image>
                </div>
              </Link>

              <Link href="/">
                <p
                  className={`text-lg ml-2 font-semibold overflow-hidden whitespace-nowrap ${
                    context.isCollapsed ? "hidden" : "block"
                  }`}
                >
                  Coffee Shop
                </p>
              </Link>
              <div
                className={`rounded-full hover:text-primary cursor-pointer p-2 ${
                  context.isCollapsed ? "m-auto" : "ml-auto mr-1"
                }`}
                onClick={context.toggleCollapse}
              >
                <RxHamburgerMenu className="w-6 h-6 " />
              </div>
            </div>

            <ul className="sidebar__list">
              {sidebarItems.map((item) => (
                <li className="sidebar__item" key={item.title}>
                  <MenuItem
                    item={item}
                    isCollapse={context.isCollapsed}
                  ></MenuItem>
                </li>
              ))}
            </ul>
          </nav>
        </aside>
      </div>
    );
}

const MenuItem = ({
  item,
  isCollapse,
}: {
  item: SidebarItem;
  isCollapse: boolean;
}) => {
  const pathname = usePathname();
  const [subMenuOpen, setSubMenuOpen] = useState(false);
  const toggleSubMenu = () => {
    setSubMenuOpen(!subMenuOpen);
  };

  return (
    <div>
      {item.submenu ? (
        <>
          <div onClick={toggleSubMenu}>
            <div
              className={`flex text-base no-underline text-black px-4 py-3 mb-2 rounded-md overflow-hidden max-h-15 hover:bg-orange-50 cursor-pointer ${
                pathname.includes(item.href) ? "bg-zinc-100" : ""
              }`}
            >
              {item.icon ? (
                <>
                  <span>
                    <item.icon className="sidebar__icon" />
                  </span>
                </>
              ) : null}

              <span
                className={`ml-2 text-lg overflow-hidden  whitespace-nowrap ${
                  isCollapse ? "hidden opacity-0" : "visible opacity-100"
                }`}
              >
                {item.title}
              </span>
              <div
                className={`ml-auto self-center ${
                  isCollapse ? "hidden" : "visible"
                } ${subMenuOpen && !isCollapse ? "rotate-180" : ""} flex`}
              >
                <ChevronDown />
              </div>
            </div>
          </div>
          {subMenuOpen && !isCollapse && (
            <div className="my-2 ml-12 flex flex-col space-y-4">
              {item.subMenuItems?.map((subItem, idx) => {
                return (
                  <Link
                    key={idx}
                    href={subItem.href}
                    className={`flex text-base no-underline text-black rounded-md overflow-hidden hover:text-primary`}
                  >
                    {subItem.icon ? (
                      <>
                        <span>
                          <subItem.icon className="sidebar__icon" />
                        </span>
                      </>
                    ) : null}

                    <span
                      className={`text-lg overflow-hidden  whitespace-nowrap ${
                        isCollapse ? "hidden" : "visible"
                      } ${
                        subItem.href === pathname
                          ? "text-primary font-medium"
                          : ""
                      }`}
                    >
                      {subItem.title}
                    </span>
                  </Link>
                );
              })}
            </div>
          )}
        </>
      ) : (
        <Link
          href={item.href}
          className={`flex text-base no-underline text-black px-4 py-3 mb-2 rounded-md overflow-hidden max-h-15 hover:bg-orange-50 ${
            item.href === pathname ? "bg-zinc-100" : ""
          }`}
        >
          {item.icon ? (
            <>
              <span>
                <item.icon className="sidebar__icon" />
              </span>
            </>
          ) : null}

          <span
            className={`ml-2 text-lg overflow-hidden  whitespace-nowrap ${
              isCollapse ? "hidden" : "visible"
            }`}
          >
            {item.title}
          </span>
        </Link>
      )}
    </div>
  );
};
