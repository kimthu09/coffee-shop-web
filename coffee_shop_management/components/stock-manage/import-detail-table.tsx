"use client";

import { CaretSortIcon } from "@radix-ui/react-icons";
import {
  ColumnDef,
  ColumnFiltersState,
  SortingState,
  VisibilityState,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable,
} from "@tanstack/react-table";

import { Button } from "@/components/ui/button";
import { Checkbox } from "@/components/ui/checkbox";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { ImportNoteDetail } from "@/types";

import { useState } from "react";
import { Input } from "../ui/input";
import { toVND } from "@/lib/utils";
import { useRouter } from "next/navigation";

export const columns: ColumnDef<ImportNoteDetail>[] = [
  {
    id: "stt",
    header: ({ table }) => (
      <div className="flex justify-center font-semibold">STT</div>
    ),
    cell: ({ row }) => (
      <div className="flex justify-center">{row.index + 1}</div>
    ),
    enableSorting: false,
    enableHiding: false,
    size: 4,
  },
  {
    accessorKey: "id",
    accessorFn: (row) => row.ingredient.id,
    header: () => {
      return <span className="font-semibold">Mã nguyên liệu</span>;
    },
    cell: ({ row }) => (
      <div className="leading-6">{row.original.ingredient.id}</div>
    ),
    size: 4,
  },
  {
    accessorKey: "name",
    header: () => {
      return <span className="font-semibold">Tên nguyên liệu</span>;
    },
    cell: ({ row }) => (
      <div className="leading-6 flex flex-col">
        {row.original.ingredient.name}
        <span className="text-muted-foreground">
          ({row.original.ingredient.measureType})
        </span>
      </div>
    ),
    size: 4,
  },
  {
    accessorKey: "amountImport",
    header: ({ column }) => (
      <div className="flex justify-end whitespace-normal">
        <span className="font-semibold">Số lượng</span>
      </div>
    ),
    cell: ({ row }) => {
      return (
        <div className="text-right font-medium">
          {row.original.amountImport.toLocaleString("vi-VN")}
        </div>
      );
    },
    size: 4,
  },
  {
    accessorKey: "price",
    header: ({ column }) => (
      <div className="flex justify-end whitespace-normal">
        <span className="font-semibold">Giá trị nhập</span>
      </div>
    ),
    cell: ({ row }) => {
      const amount = parseFloat(row.getValue("price"));

      return (
        <div className="text-right font-medium">
          {toVND(amount)}
          {/* <div className="text-sm flex font-light items-center justify-end gap-1">
            {row.original.supplier.name} <BiBox className="h-4 w-4" />
          </div> */}
        </div>
      );
    },
    size: 4,
  },
  {
    accessorKey: "totalUnit",
    accessorFn: (row) => row.amountImport * row.price,
    header: ({ column }) => (
      <div className="flex justify-end whitespace-normal">
        <span className="font-semibold">Thành tiền</span>
      </div>
    ),
    cell: ({ row }) => {
      return (
        <div className="text-right font-medium">
          {toVND(row.original.amountImport * row.original.price)}
        </div>
      );
    },
    size: 4,
  },
];
export function ImportDetailTable(details: ImportNoteDetail[]) {
  const data = Object.values(details);

  const router = useRouter();
  const [sorting, setSorting] = useState<SortingState>([]);
  const [columnFilters, setColumnFilters] = useState<ColumnFiltersState>([]);
  const [columnVisibility, setColumnVisibility] = useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = useState({});
  const table = useReactTable({
    data,
    columns,
    onSortingChange: setSorting,
    onColumnFiltersChange: setColumnFilters,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    onRowSelectionChange: setRowSelection,
    state: {
      sorting,
      columnFilters,
      columnVisibility,
      rowSelection,
    },
    initialState: {
      pagination: {
        pageSize: 1000,
      },
    },
  });
  return (
    <div className="rounded-md border w-full">
      <Table>
        <TableHeader>
          {table.getHeaderGroups().map((headerGroup) => (
            <TableRow
              key={headerGroup.id}
              className="bg-orange-50 hover:bg-orange-50"
            >
              {headerGroup.headers.map((header) => {
                return (
                  <TableHead key={header.id}>
                    {header.isPlaceholder
                      ? null
                      : flexRender(
                          header.column.columnDef.header,
                          header.getContext()
                        )}
                  </TableHead>
                );
              })}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {table.getRowModel().rows?.length ? (
            table.getRowModel().rows.map((row, index) => (
              <TableRow
                key={row.id}
                data-state={row.getIsSelected() && "selected"}
              >
                {row.getVisibleCells().map((cell) => (
                  <TableCell key={cell.id}>
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columns.length} className="h-24 text-center">
                No results.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}
