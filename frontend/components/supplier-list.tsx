import { useState } from "react";
import Loading from "./loading";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "./ui/dropdown-menu";
import { Button } from "./ui/button";
import { LuCheck, LuChevronsUpDown } from "react-icons/lu";
import {
  Command,
  CommandEmpty,
  CommandGroup,
  CommandInput,
  CommandItem,
} from "./ui/command";
import getAllSupplier from "@/lib/getAllSupplier";
import { cn } from "@/lib/utils";
import { FaPlus } from "react-icons/fa";
import { get } from "http";
import { getToken } from "@/lib/auth";

export interface SupplierListProps {
  supplierId: string;
  setSupplierId: (id: string) => void;
  canAdd?: boolean;
}
const SupplierList = ({
  supplierId,
  setSupplierId,
  canAdd,
}: SupplierListProps) => {
  const [open, setOpen] = useState(false);
  const token = getToken();
  const { suppliers, isLoading, isError, mutate } = getAllSupplier(token!);
  const handlePublisherAdded = async (publisherId: string) => {
    await mutate();
    //TODO
    // setPublisherId(publisherId);
  };
  if (isError) return <div>Failed to load</div>;
  if (!suppliers) {
    <Loading />;
  } else
    return (
      <div className="flex gap-1">
        <DropdownMenu open={open} onOpenChange={setOpen}>
          <DropdownMenuTrigger asChild>
            <Button
              variant="outline"
              role="combobox"
              aria-expanded={open}
              className="justify-between w-full pl-2"
            >
              {supplierId
                ? suppliers.find((item: any) => item.id === supplierId)?.name
                : "Chọn nhà cung cấp"}
              <LuChevronsUpDown className="ml-1 h-4 w-4 shrink-0 opacity-50" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="DropdownMenuContent">
            <Command>
              <CommandInput
                placeholder="Tìm số điện thoại nhà cung cấp"
                // onValueChange={(str) => setNewCategory(str)}
              />
              <CommandEmpty className="py-2 px-6">
                <div className="text-sm">Không tìm thấy nhà cung cấp</div>
              </CommandEmpty>
              <CommandGroup className="max-h-48 overflow-y-auto">
                {suppliers.map((item: any) => (
                  <CommandItem
                    value={item.phone}
                    key={item.id}
                    onSelect={() => {
                      setSupplierId(item.id);
                      setOpen(false);
                    }}
                  >
                    <LuCheck
                      className={cn(
                        "mr-2 h-4 w-4",
                        item.id === supplierId ? "opacity-100" : "opacity-0"
                      )}
                    />
                    <div className="flex flex-col">
                      {item.name}
                      <span className="text-muted-foreground">
                        {item.phone}
                      </span>
                    </div>
                  </CommandItem>
                ))}
              </CommandGroup>
            </Command>
          </DropdownMenuContent>
        </DropdownMenu>
        {/* {canAdd ? (
          <CreatePublisher handlePublisherAdded={handlePublisherAdded}>
            <Button type="button" size={"icon"} className="px-3">
              <FaPlus />
            </Button>
          </CreatePublisher>
        ) : null} */}
      </div>
    );
};

export default SupplierList;
