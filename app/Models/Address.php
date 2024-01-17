<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use App\Models\customerStore;
use Illuminate\Http\Request;
use Illuminate\Database\Eloquent\SoftDeletes;
class Address extends Model
{
    use HasFactory, SoftDeletes;
    protected $table = 'address';

    protected $fillable = [
        'id',
        'customer_store_id',
        'address',
        'district',
        'city',
        'province',
        'postal_code',
    ];
    protected $hidden = [
        'deleted_at'
    ];

    public function customerStore()
    {
        return $this->belongsTo(CustomerStore::class);
    }

    public static function simpan($validatedData, $id)
    
    {
        // Membuat instance address baru 
        $address = new Address($validatedData);

        $address->customer_store_id = $id;
        
        $address->updated_at = null;

        $address->created_at = now();
        


        // menyimpan Address
        $address->save();

        return  $address;
    }
}

