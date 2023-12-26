import { StaffListProps } from "@/types";
import { useEffect, useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import { Command, CommandGroup, CommandInput, CommandItem } from "./ui/command";
import { LuCheck, LuChevronsUpDown } from "react-icons/lu";
import { Button } from "./ui/button";
import { cn } from "@/lib/utils";
import Loading from "./loading";
import getAllStaff from "@/lib/getAllStaffClient";
import { getToken } from "@/lib/auth";

const StaffList = ({ staff, setStaff }: StaffListProps) => {
  const [openRole, setOpenRole] = useState(false);
  const token = getToken();
  const { staffs, isLoading, isError } = getAllStaff(token!);
  useEffect(() => {
    // handleStaffSelected(staff);
  }, [staff]);
  if (isError) return <div>Failed to load</div>;
  if (!staffs) {
    <Loading />;
  } else
    return (
      <DropdownMenu open={openRole} onOpenChange={setOpenRole}>
        <DropdownMenuTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={openRole}
            className="justify-between flex-1 min-w-0"
          >
            {staff
              ? staffs.find((item) => item.id === staff)?.name
              : "Chọn nhân viên"}
            <LuChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent className="DropdownMenuContent">
          <Command>
            <CommandInput
              placeholder="Tìm tên nhân viên"
              // onValueChange={(str) => setNewCategory(str)}
            />
            <CommandGroup className="p-0">
              {staffs.map((item) => (
                <CommandItem
                  value={item.name}
                  key={item.id}
                  onSelect={() => {
                    // handleSelected(item.id);
                    setStaff(item.id);
                    setOpenRole(false);
                  }}
                >
                  <LuCheck
                    className={cn(
                      "mr-1 h-4 w-4",
                      item.id === staff ? "opacity-100" : "opacity-0"
                    )}
                  />
                  <div className="flex flex-col">
                    {item.name}
                    <span className=" text-muted-foreground">{item.id}</span>
                  </div>
                </CommandItem>
              ))}
            </CommandGroup>
          </Command>
        </DropdownMenuContent>
      </DropdownMenu>
    );
};

export default StaffList;
