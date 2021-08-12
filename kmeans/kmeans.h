#ifndef KMEANS
#define KMEANS
/*
* represent one pixel in image, with attribute r, g and b;
* we do not care about its position here 
*/
typedef struct{
    int r;
    int g;
    int b;
    /* to specify which center this pixel belongs to,
    * client do not need to set or get this attribute. 
     */
    int clusterId;
} Pixel;

/*
 * Function: kmeans
 * Description: do kmeans clustering 
 * Input: src. pointer to an array of Pixel, which provide original image pixes information.
 * Input: dst. pointer to an array of Pixel, which received pixes information about clustering.
 * Input: pixesLen. size of src and dst 
 * Input: k. cluster count
 * Return: 0 means success
 */
int kmeans(Pixel* src, Pixel* dst, int pixesLen, int k);

#endif
