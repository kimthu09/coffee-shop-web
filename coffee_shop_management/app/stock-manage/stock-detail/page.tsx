import IngredientDetailsTable from "@/components/stock-manage/ingredient-detail-table";
import { Card, CardContent } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { ingredientDetails, ingredients } from "@/constants";

const StockDetail = ({ searchParams }: { searchParams: { id: string } }) => {
  const ingredient = ingredients.find((item) => item.id === searchParams.id);
  const details = ingredientDetails.filter(
    (item) => item.idIngre === searchParams.id
  );
  return (
    <div className="col items-center">
      <div className="col xl:w-4/5 w-full xl:px-0 md:px-8 px-0">
        <h1 className="xl:text-3xl text-2xl">Chi tiết nguyên liệu</h1>
        <Card>
          <CardContent className="p-6 flex flex-col   gap-4">
            <div className="flex gap-4 lg:flex-row flex-col">
              <div className="basis-1/3">
                <Label htmlFor="id">Mã nguyên liệu</Label>
                <Input id="id" value={ingredient?.id} readOnly></Input>
              </div>
              <div className="basis-2/3">
                <Label htmlFor="name">Tên nguyên liệu</Label>
                <Input id="name" value={ingredient?.name} readOnly></Input>
              </div>
            </div>
            <div className="flex flex-col gap-4 lg:flex-row">
              <div className="flex gap-4 lg:basis-2/3 sm:flex-row flex-col">
                <div className="flex-1">
                  <Label htmlFor="sl">Tồn cuối</Label>
                  <Input id="sl" value={ingredient?.total} readOnly></Input>
                </div>
                <div className="flex-1">
                  <Label htmlFor="dv">Đơn vị</Label>
                  <Input id="dv" value={ingredient?.unit.name} readOnly></Input>
                </div>
              </div>
              <div className="self-end sm:w-1/2 w-full sm:pl-2 lg:basis-1/3 lg:pl-0">
                <Label htmlFor="gia">Giá</Label>
                <Input id="gia" value={ingredient?.price} readOnly></Input>
              </div>
            </div>
          </CardContent>
        </Card>

        <div className="my-4 p-3 sha bg-white shadow-[0_1px_3px_0_rgba(0,0,0,0.2)]">
          <IngredientDetailsTable
            id={ingredient?.id!}
            ingredientDetails={details}
          />
        </div>
      </div>
    </div>
  );
};

export default StockDetail;
