import { UnitListProps } from "@/types";
import { useState } from "react";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "./ui/command";
import { Check, ChevronsUpDown } from "lucide-react";
import { Button } from "./ui/button";
import { measureUnits } from "@/constants";
import { cn } from "@/lib/utils";

const UnitList = ({ unit, setUnit }: UnitListProps) => {
  const [openUnit, setOpenUnit] = useState(false);
  return (
    <DropdownMenu open={openUnit} onOpenChange={setOpenUnit}>
      <DropdownMenuTrigger asChild>
        <Button
          id="cateList"
          variant="outline"
          role="combobox"
          aria-expanded={openUnit}
          className="justify-between w-full"
        >
          {unit
            ? measureUnits.find((item) => item.name === unit)?.name
            : "Chọn đơn vị"}
          <ChevronsUpDown className="ml-2 h-4 w-4 shrink-0 opacity-50" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent className=" DropdownMenuContent">
        <Command>
          <CommandInput placeholder="Tìm điều kiện lọc" />
          <CommandEmpty className="p-2 text-sm">
            Không tìm thấy điều kiện lọc.
          </CommandEmpty>
          <CommandGroup>
            {measureUnits.map((item) => (
              <CommandItem
                value={item.name}
                key={item.id}
                onSelect={() => {
                  setUnit(item.name);
                  setOpenUnit(false);
                }}
              >
                <Check
                  className={cn(
                    "mr-2 h-4 w-4",
                    item.name === unit ? "opacity-100" : "opacity-0"
                  )}
                />
                {item.name}
              </CommandItem>
            ))}
            <CommandItem
              key={""}
              onSelect={() => {
                setUnit("");
                setOpenUnit(false);
              }}
            >
              <Check
                className={cn(
                  "mr-2 h-4 w-4",
                  "" === unit ? "opacity-100" : "opacity-0"
                )}
              />
              {"Chọn đơn vị"}
            </CommandItem>
          </CommandGroup>
        </Command>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export default UnitList;
