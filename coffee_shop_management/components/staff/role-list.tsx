import { RoleListProps } from "@/types";
import { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "../ui/command";
import { Check, ChevronsUpDown } from "lucide-react";
import { Button } from "../ui/button";
import { cn } from "@/lib/utils";
import { roles } from "@/constants";

const RoleList = ({ role, setRole }: RoleListProps) => {
  const [openRole, setOpenRole] = useState(false);
  return (
    <DropdownMenu open={openRole} onOpenChange={setOpenRole}>
      <DropdownMenuTrigger asChild>
        <Button
          id="cateList"
          variant="outline"
          role="combobox"
          aria-expanded={openRole}
          className="justify-between w-full"
        >
          {role
            ? roles.find((item) => item.name === role)?.name
            : "Chọn phân quyền"}
          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className=" DropdownMenuContent">
        <Command>
          <CommandGroup>
            {roles.map((item) => (
              <CommandItem
                value={item.name}
                key={item.id}
                onSelect={() => {
                  setRole(item.name);
                  setOpenRole(false);
                }}
              >
                <Check
                  className={cn(
                    "mr-2 h-4 w-4",
                    item.name === role ? "opacity-100" : "opacity-0"
                  )}
                />
                {item.name}
              </CommandItem>
            ))}
            <CommandItem
              key={""}
              onSelect={() => {
                setRole("");
                setOpenRole(false);
              }}
            >
              <Check
                className={cn(
                  "mr-2 h-4 w-4",
                  "" === role ? "opacity-100" : "opacity-0"
                )}
              />
              {"Chọn phân quyền"}
            </CommandItem>
          </CommandGroup>
        </Command>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default RoleList;
