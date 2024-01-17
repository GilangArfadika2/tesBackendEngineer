<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use Illuminate\Http\Request;
use App\Models\Address;

use Illuminate\Database\Eloquent\SoftDeletes;
// use Illuminate\Database\Eloquent\SoftDeletes;

//  Model Customer Store
class CustomerStore extends Model
{
    
    use HasFactory, SoftDeletes;
    protected $table = 'customer_stores';

    protected $fillable = [
        'id',
        'title',
        'name',
        'gender',
        'phone_number',
        'image',
        'email'
    ];

    protected $hidden = [
        'created_at', 'updated_at','deleted_at'
    ];

    public function address()
    {
        return $this->hasMany(Address::class);
    }
    // Menyimmpan data customer store baru
    public static function simpan($validatedData)
    
    {  
        
        $customerStore = new CustomerStore($validatedData);
        

        // set updated_at dengan nilai null
        $customerStore->updated_at = null;
        
        // set created_at dengan tanggal sekarang
        $customerStore->created_at = now();

        // simpan ke database
        $customerStore->save();

        return $customerStore;
    }

   

    
}
