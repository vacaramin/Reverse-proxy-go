<?php 



$agent = 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36';
$Uri = '/app/documents';
$post = file_get_contents('php://input');
$r = ['accept', 'accept-language', 'x-requested-with', 'main-request', 'x-newrelic-id', 'x-xsrf-token', 'cache-control', 'content-type', 'content-length', 'authorization', 'x-access-token', 'x-human-token', 'x-csrf-token', 'x-requested-with'];



if ($_SERVER['REQUEST_URI'] == '/')
{
    header('location: ' . $Uri);
    exit();
}
else
{
    $q = 'https://beta.frase.io' . $_SERVER['REQUEST_URI'];
}
foreach (getallheaders() as $n => $v)
{
    if (in_array(strtolower($n) , $r))
    {
        $headers[] = $n . ':' . $v;
    }
}
$headers[] = 'Cookie:' . $cookie;
$headers[] = 'Origin: https://beta.frase.io';
function secret($curl, $bc)
{
    $r2 = ['set-cookie', 'content-length', 'transfer-encoding'];
    $king = 0;
    foreach ($r2 as $v)
    {
        if (strlen($bc) > 0 && strpos(strtolower($bc) , strtolower($v)) === false)
        {
            $king++;
        }
    }
    if ($king == sizeof($r2) && strlen($bc) > 2)
    {
        header($bc);
    }
    return strlen($bc);
}

    $curl = curl_init();
    curl_setopt($curl, CURLOPT_URL, $q);
    curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);
    curl_setopt($curl, CURLOPT_FOLLOWLOCATION, 1);
    curl_setopt($curl, CURLOPT_HTTPHEADER, $headers);
    curl_setopt($curl, CURLOPT_USERAGENT, $agent);
    //curl_setopt($curl, CURLOPT_COOKIEFILE, $cookie);
    curl_setopt($curl, CURLOPT_HEADERFUNCTION, 'secret');
    curl_setopt($curl, CURLOPT_CUSTOMREQUEST, $_SERVER['REQUEST_METHOD']);
    curl_setopt($curl, CURLOPT_POSTFIELDS, $post);
    $result = curl_exec($curl);
    $result  = str_replace("</body>", $replaceCode, $result);
    $result = str_replace('beta.frase.io', 'spyfu.host', $result);
   $result = str_replace('</head>', '<style>.pt-3{display:none!important;}</style></head>', $result);

    
   


    $contentTypeOrg = curl_getinfo($curl, CURLINFO_CONTENT_TYPE);
    $contentType = strstr($contentTypeOrg, ';', true);
    curl_close($curl);
    echo $result;
    exit();
    header('Content-Type: ' . $contentTypeOrg);
    


?>